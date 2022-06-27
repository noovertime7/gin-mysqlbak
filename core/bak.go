package core

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/public/ding"
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
	DingConfig   *dao.DingDatabase
	OssConfig    *dao.OssDatabase
	DingStatus   int
	OssStatus    int
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
		DingConfig:   detail.Ding,
		OssConfig:    detail.Oss,
		AfterBakChan: afterBakChan,
	}
	en, err := xorm.NewEngine("mysql", bakhandler.User+":"+bakhandler.PassWord+"@tcp("+bakhandler.Host+")/"+bakhandler.DbName+"?charset=utf8&parseTime=true")
	if err != nil {
		log.Logger.Errorf("创建数据库连接失败:%s", err)
		return nil, err
	}
	bakhandler.Engine = en
	c := cron.New()
	if _, ok := CronJob[bakhandler.TaskID]; ok {
		return nil, errors.New("当前备份任务已启动，切勿重复启动")
	}
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
	_, err := b.Cron.AddJob(b.BackupCycle, b)
	if err != nil {
		return err
	}
	log.Logger.Infof("创建备份任务成功，任务id：%d", b.TaskID)
	// 备份前准备工作
	b.BeforBak()
	//启动数据库备份服务
	b.Cron.Start()
	return nil
}

func (b *BakHandler) IsStart(tid int) bool {
	if _, ok := CronJob[tid]; !ok {
		return false
	}
	return true
}

func (b *BakHandler) StopBak(tid int) error {
	log.Logger.Debug("StopBak", CronJob)
	for id, cronjob := range CronJob {
		if id == tid {
			cronjob.Stop()
			delete(CronJob, tid)
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
		b.BakStatus = 0
		b.OssStatus = 2
		b.DingStatus = 2
		b.BakMsg = fmt.Sprintf("%s", err)
		b.FileName = "unknown"
		AfterBak(b)
		log.Logger.Error("备份失败", err)
		return
	}

	b.BakMsg = "success"
	b.BakStatus = 1
	b.FileSize = public.GetFileSize(b.FileName)
	//首先判定钉钉 oss都操作成功，状态改为1
	b.DingStatus = 2
	b.OssStatus = 2
	//判断是否启动钉钉提醒
	AfterBak(b)
	//发送对象到channel
	log.Logger.Info("备份数据库成功,保存备份历史到数据库,发送消息")
	b.AfterBakChan <- b
	log.Logger.Debug("BakHandler 发送消息成功")
}

func AfterBak(b *BakHandler) {
	//判断是否启动钉钉提醒
	if b.DingConfig.IsDingSend == 1 {
		baktime := time.Now().Format("2006年01月02日15:04:01")
		markcontent := map[string]string{
			"title": b.Host + b.DbName + "备份状态",
			"text": fmt.Sprintf(
				"\n- 备份时间:%v\n- 备份状态:%s\n- OSS上传状态:%s\n- 备份文件目录:%s\n![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n", baktime, b.BakMsg, b.OssStatus, b.FileName),
		}
		webhook := ding.Webhook{AtAll: true, Secret: b.DingConfig.DingSecret, AccessToken: b.DingConfig.DingAccessToken}
		log.Logger.Infof("%s:%s开始发送钉钉消息", b.Host, b.DbName)
		if err := webhook.SendMarkDown(markcontent); err != nil {
			b.DingStatus = 0
			log.Logger.Error("钉钉消息发送失败", err)
			return
		}
	}
	b.DingStatus = 1
	log.Logger.Infof("%s:%s发送钉钉消息成功", b.Host, b.DbName)
	//判断是否启动OSS保存
	if b.OssConfig.IsOssSave == 1 {

	}
}
