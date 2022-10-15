package system

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdao"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/agentservice"
	"github.com/noovertime7/gin-mysqlbak/agent/repository"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/public/database"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/robfig/cron/v3"
	"strings"
	"sync"
	"time"
)

type dateNumInfoSync struct {
	ctx       context.Context
	c         *cron.Cron
	lock      sync.RWMutex
	Once      sync.Once
	ReSync    string
	StoreInfo chan *agentdao.AgentDateInfo
}

func NewDateNumInfoSync(ctx context.Context, ReSync string) *dateNumInfoSync {
	return &dateNumInfoSync{
		ctx:       ctx,
		c:         cron.New(),
		lock:      sync.RWMutex{},
		Once:      sync.Once{},
		ReSync:    ReSync,
		StoreInfo: make(chan *agentdao.AgentDateInfo, 100),
	}
}

func (d *dateNumInfoSync) Start() {
	EntryID, err := d.c.AddJob(d.ReSync, d)
	if err != nil {
		log.Logger.Errorf("日期信息同步定时任务启动任务失败,任务ID:%d", EntryID)
		d.handleErr(err)
		return
	}
	d.c.Start()
	log.Logger.Infof("日期信息同步定时任务启动任务成功,任务ID:%d", EntryID)
}

func (d *dateNumInfoSync) Run() {
	d.Once.Do(func() {
		go d.Store(d.ctx)
	})
	var affect int
	log.Logger.Info("启动日期信息定时器，开始循环日期信息数据到服务端")
	ts := agentservice.GetClusterTaskService()
	dates := public.GetBeforeDates(15)
	//1、查询当前所有服务
	var AgentService *repository.AgentService
	for _, date := range dates {
		log.Logger.Infof("当前查询日期%s", date)
		var taskNum int64
		var finNum int64
		services, err := AgentService.GetHealthAgentList(d.ctx, &agentdto.AgentListInput{PageNo: 1, PageSize: 9999})
		if err != nil {
			d.handleErr(err)
			return
		}
		for _, service := range services.AgentOutPutItem {
			info, err := ts.GetDateNumInfo(d.ctx, &agentdto.GetDateNumInfoInput{Date: date, ServiceName: service.ServiceName})
			if err != nil {
				d.handleErr(err)
				return
			}
			log.Logger.Infof("从客户端%s获取到数据%v", service.ServiceName, info)
			if info == nil {
				log.Logger.Errorf("当前服务%s调用失败，调用地址:%s,请检查", service.ServiceName, service.Address)
				return
			}
			taskNum += info.TaskNum
			finNum += info.FinishNum
		}
		dateDB := &agentdao.AgentDateInfo{CurrentTime: date, TaskNum: taskNum, FinishNum: finNum}
		d.StoreInfo <- dateDB
		affect++
	}
	d.storeJobHistory(d.ctx, affect)
}

func (d *dateNumInfoSync) Store(ctx context.Context) {
	tx := database.GetDB()
	for {
		select {
		case dateDB := <-d.StoreInfo:
			//先去查询一下看看有没有
			tempDb := &agentdao.AgentDateInfo{CurrentTime: dateDB.CurrentTime}
			dateinfo, err := tempDb.Find(d.ctx, tx, tempDb)
			if err != nil {
				return
			}
			if dateinfo.Id != 0 {
				dateinfo.TaskNum = dateDB.TaskNum
				dateinfo.FinishNum = dateDB.FinishNum
				dateinfo.CurrentTime = dateDB.CurrentTime
				if err := dateinfo.Updates(d.ctx, tx); err != nil {
					d.handleErr(err)
					return
				}
			} else {
				dateDB := &agentdao.AgentDateInfo{
					TaskNum:     dateDB.TaskNum,
					FinishNum:   dateDB.FinishNum,
					CurrentTime: dateDB.CurrentTime,
				}
				if err := dateDB.Save(d.ctx, tx); err != nil {
					d.handleErr(err)
					return
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

func (d *dateNumInfoSync) storeJobHistory(ctx context.Context, affected int) {
	jobHistoryDB := &dao.JobHistory{
		JobType:    public.DateNumInfoJob,
		JobCycle:   d.ReSync,
		Affected:   affected,
		UpdateTime: time.Now(),
	}
	errs, ok := d.GetErr()
	if !ok {
		jobHistoryDB.Status = sql.NullInt64{Int64: 0, Valid: true}
		jobHistoryDB.Message = fmt.Sprintf(strings.Join(errs, ";"))
		success = true
		errMessage = []string{}
	} else {
		jobHistoryDB.Status = sql.NullInt64{Int64: 1, Valid: true}
		jobHistoryDB.Message = "集群任务同步定时任务执行成功"
	}
	j, err := jobHistoryDB.Find(ctx, database.GetDB(), &dao.JobHistory{JobType: jobHistoryDB.JobType})
	if err != nil {
		jobHistoryDB.Status = sql.NullInt64{Int64: 0, Valid: true}
		jobHistoryDB.Message = err.Error()
	}
	if j.ID != 0 {
		jobHistoryDB.ID = j.ID
		if err := jobHistoryDB.Updates(ctx, database.GetDB()); err != nil {
			log.Logger.Error("集群任务同步定时任务保存JobHistory数据库失败")
			return
		}
	} else {
		if err := jobHistoryDB.Save(ctx, database.GetDB()); err != nil {
			log.Logger.Error("集群任务同步定时任务保存JobHistory数据库失败")
			return
		}
	}
}

var (
	dateSuccess    = true
	dateErrMessage []string
)

func (d *dateNumInfoSync) handleErr(err error) {
	dateSuccess = false
	dateErrMessage = append(dateErrMessage, err.Error())
}

func (d *dateNumInfoSync) GetErr() ([]string, bool) {
	return dateErrMessage, dateSuccess
}

func (d *dateNumInfoSync) Stop() {
	d.c.Stop()
}
