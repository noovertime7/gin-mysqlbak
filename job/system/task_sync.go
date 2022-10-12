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

type taskSync struct {
	ctx       context.Context
	c         *cron.Cron
	lock      sync.RWMutex
	ReSync    string
	StoreInfo chan *agentdao.TaskOverview
}

func NewTaskSyncJob(ctx context.Context, ReSync string) *taskSync {
	return &taskSync{
		ctx:       ctx,
		c:         cron.New(),
		lock:      sync.RWMutex{},
		ReSync:    ReSync,
		StoreInfo: make(chan *agentdao.TaskOverview, 100),
	}
}

func (t *taskSync) Start() {
	EntryID, err := t.c.AddJob(t.ReSync, t)
	if err != nil {
		log.Logger.Errorf("任务同步定时任务启动任务失败,任务ID:%d", EntryID)
		return
	}
	go t.Store(t.ctx)
	t.c.Start()
	log.Logger.Infof("任务同步定时任务启动任务成功,任务ID:%d", EntryID)
}

func (t *taskSync) Run() {
	var affect int
	log.Logger.Info("启动集群全局任务管理定时器，开始循环同步任务数据到服务端")
	h := agentservice.GetClusterHostService()
	ts := agentservice.GetClusterTaskService()
	//1、查询当前所有服务
	var AgentService *repository.AgentService
	services, err := AgentService.GetAgentList(t.ctx, &agentdto.AgentListInput{PageNo: 1, PageSize: 9999})
	if err != nil {
		log.Logger.Error(err)
	}
	for _, service := range services.AgentOutPutItem {
		//2、通过服务循环查询所有主机
		hosts, err := h.HostList(t.ctx, &agentdto.HostListInput{
			ServiceName: service.ServiceName,
			Info:        "",
			PageNo:      1,
			PageSize:    9999,
		})
		if err != nil {
			t.handleErr(err)
			return
		}
		for _, host := range hosts.ListItem {
			//3、通过主机循环查询主机下所有任务
			tasks, err := ts.GetTaskUnscopedList(t.ctx, &agentdto.TaskListInput{
				ServiceName: service.ServiceName,
				HostId:      host.ID,
				Info:        "",
				PageNo:      1,
				PageSize:    9999,
			})
			if err != nil {
				t.handleErr(err)
				return
			}
			for _, task := range tasks.TaskListItem {
				//4、到保存数据库
				taskOverDB := &agentdao.TaskOverview{
					ServiceName: service.ServiceName,
					HostId:      host.ID,
					Host:        host.Host,
					Type:        host.Type,
					TaskId:      task.ID,
					DbName:      task.DBName,
					BackupCycle: task.BackupCycle,
					KeepNumber:  task.KeepNumber,
					FinishNum:   task.FinishNum,
					Status:      sql.NullInt64{Int64: task.Status, Valid: true},
					CreatedAt:   public.StringToTime(task.CreateAt),
					UpdateAt:    public.StringToTime(task.UpdateAt),
					IsDeleted:   sql.NullInt64{Int64: task.IsDeleted, Valid: true},
					DeletedAt:   public.StringToTime(task.DeletedAt),
				}
				t.StoreInfo <- taskOverDB
				affect++
			}
		}
	}
	t.storeJobHistory(t.ctx, affect)
	return
}

var (
	success    = true
	errMessage []string
)

func (t *taskSync) handleErr(err error) {
	success = false
	errMessage = append(errMessage, err.Error())
}

func (t *taskSync) GetErr() ([]string, bool) {
	return errMessage, success
}

func (t *taskSync) Stop() {
	t.c.Stop()
}

func (t *taskSync) storeJobHistory(ctx context.Context, affected int) {
	jobHistoryDB := &dao.JobHistory{
		JobType:    public.TaskSyncJob,
		JobCycle:   t.ReSync,
		Affected:   affected,
		UpdateTime: time.Now(),
	}
	errs, ok := t.GetErr()
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

func (t *taskSync) Store(ctx context.Context) {
	tx := database.GetDB()
	for {
		select {
		case taskOverDB := <-t.StoreInfo:
			//先去查询一下看看有没有
			temp, err := taskOverDB.Find(ctx, tx, &agentdao.TaskOverview{ServiceName: taskOverDB.ServiceName, TaskId: taskOverDB.TaskId, HostId: taskOverDB.HostId})
			if err != nil {
				t.handleErr(err)
				return
			}
			if temp.ID != 0 {
				//更新操作
				log.Logger.Debug("任务总览定时器开始更新数据库", temp)
				if err := taskOverDB.Updates(ctx, tx, temp.ID); err != nil {
					t.handleErr(err)
					return
				}
			} else {
				//没有查到，新增操作
				log.Logger.Debug("任务总览定时器开始插入数据库", taskOverDB)
				if err := taskOverDB.Save(ctx, tx); err != nil {
					t.handleErr(err)
					return
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
