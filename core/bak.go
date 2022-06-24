package core

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"os"
	"strconv"
	"strings"
	"time"
)

type BakHandler struct {
	TaskID       int
	Cron         *cron.Cron
	Engine       *xorm.Engine
	Host         string
	BackInfoId   string
	PassWord     string
	User         string
	Port         string
	DbName       string
	BackupCycle  string
	KeepNumber   int
	ISAllDBBak   int
	ISDingSender int
	ISOssSave    int
	BakStatus    int
	BakMsg       string
	FileName     string
	FileSize     int
	AfterBakChan chan *BakHandler
}

var CronJob = make(map[int]*cron.Cron)

func NewBakController(detail *dao.TaskDetail, afterBakChan chan *BakHandler) (*BakHandler, error) {
	bakhandler := &BakHandler{
		TaskID:       detail.Info.Id,
		Host:         detail.Info.Host,
		PassWord:     detail.Info.Password,
		User:         detail.Info.User,
		DbName:       detail.Info.DBName,
		BackupCycle:  detail.Info.BackupCycle,
		KeepNumber:   detail.Info.KeepNumber,
		ISAllDBBak:   detail.Info.IsAllDBBak,
		ISDingSender: detail.Ding.IsDingSend,
		ISOssSave:    detail.Oss.IsOssSave,
		AfterBakChan: afterBakChan,
	}
	en, err := xorm.NewEngine("mysql", bakhandler.User+":"+bakhandler.PassWord+"@tcp("+bakhandler.Host+")/"+bakhandler.DbName+"?charset=utf8&parseTime=true")
	if err != nil {
		log.Logger.Errorf("创建数据库连接失败:%s", err)
		return nil, err
	}
	bakhandler.Engine = en
	c := cron.New()
	jobmap := make(map[int]*cron.Cron)
	jobmap[bakhandler.TaskID] = c
	CronJob = jobmap
	bakhandler.Cron = c
	log.Logger.Debug("NewBakHandler", CronJob)
	return bakhandler, nil
}

func (b *BakHandler) BeforBak() {
	path, _ := os.Getwd()
	public.CreateDir(path + "/bakfile/")
	dir := path + "/bakfile/" + strconv.Itoa(b.TaskID)
	baktime := time.Now().Format("2006-01-02-15-04")
	public.CreateDir(dir)
	host := strings.Split(b.Host, ":")[0]
	file := dir + "/" + host + "_" + b.DbName + "_" + baktime + ".sql"
	b.FileName = file
}

func (b *BakHandler) StartBak() error {
	eid, err := b.Cron.AddJob(b.BackupCycle, b)
	if err != nil {
		return err
	}
	log.Logger.Infof("创建备份任务成功，任务id：%d", eid)
	// 备份前准备工作
	b.BeforBak()
	//启动数据库备份服务
	b.Cron.Start()
	return nil
}

func (b *BakHandler) StopBak(tid int) error {
	log.Logger.Debug("StopBak", CronJob)
	if _, ok := CronJob[tid]; !ok {
		return errors.New("备份任务不存在请检查备份id")
	}
	for id, cronjob := range CronJob {
		if id == tid {
			cronjob.Stop()
			log.Logger.Infof("任务ID:%d,备份库名:%s 停止成功", id, b.DbName)
			return nil
		}
	}

	return nil
}

func (b *BakHandler) Run() {
	log.Logger.Info("BakHandler 开始备份数据库")
	err := b.Engine.DumpAllToFile(b.FileName)
	if err != nil {
		b.BakStatus = 1
		b.BakMsg = fmt.Sprintf("%s", err)
		b.FileName = ""
		log.Logger.Error("备份失败", err)
		return
	}
	b.FileSize = public.GetFileSize(b.FileName)
	b.BakStatus = 0
	b.BakMsg = "success"
	log.Logger.Info("备份数据库成功,保存备份历史到数据库,发送消息")
	b.AfterBakChan <- b
	log.Logger.Debug("BakHandler 发送消息成功")
}
