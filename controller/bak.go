package controller

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/core"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/public/database"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type BakController struct {
	AfterBakChan chan *core.BakHandler
}

func BakRegister(group *gin.RouterGroup) {
	bak := &BakController{}
	group.GET("/start", bak.StartBak)
	group.GET("/start_bak_all", bak.StartBakAll)
	group.GET("/start_bak_all_byhost", bak.StartBakAllByHost)
	group.GET("/stop", bak.StopBak)
	group.GET("/stop_bak_all_byhost", bak.StopBakAllByHost)
	group.GET("/stop_bak_all", bak.StopBakAll)
}

func (bak *BakController) StartBak(c *gin.Context) {
	tx := database.GetDB()
	params := &dto.Bak{}
	if err := params.BindValidParams(c); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	taskinfo := &dao.TaskInfo{Id: params.ID}
	taskdetail, err := taskinfo.TaskDetail(c, tx, taskinfo)
	if err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 2002, err)
		return
	}
	tx, err = lib.GetGormPool("default")
	if err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 2002, err)
		return
	}
	bak.AfterBakChan = make(chan *core.BakHandler, 10)
	go bak.ListenAndSave(c, database.GetDB(), bak.AfterBakChan)
	bakhandler, err := core.NewBakHandler(taskdetail, bak.AfterBakChan)
	if err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 2003, err)
		return
	}
	if err = bakhandler.StartBak(); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 2003, err)
		return
	}
	taskinfo.Status = 1
	if err = taskinfo.Updates(c, database.GetDB()); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 2004, err)
		return
	}
	middleware.ResponseSuccess(c, "启动任务成功")
}

func (bak *BakController) StartBakAll(c *gin.Context) {
	var b BakController
	taskinfo := &dao.TaskInfo{}
	result, err := taskinfo.FindAllTask(c, database.GetDB(), nil)
	if err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 2002, err)
		return
	}
	log.Logger.Debug("StartBakAll 任务查询结果", result)
	if len(result) == 0 {
		log.Logger.Info("当前主机备份任务为空")
		middleware.ResponseError(c, 2003, errors.New("当前主机备份任务为空"))
		return
	}
	b.AfterBakChan = make(chan *core.BakHandler, 10)
	tx := database.GetDB()
	go b.ListenAndSave(c, tx, b.AfterBakChan)
	for _, task := range result {
		if task.Status == 1 {
			log.Logger.Infof("TASK_ID:%d,HOSTID:%d,%s数据库任务已经开启,返回", task.Id, task.HostID, task.DBName)
			continue
		}
		taskdetail, err := taskinfo.TaskDetail(c, tx, task)
		if err != nil {
			log.Logger.Error(err)
			return
		}
		bakhandler, err := core.NewBakHandler(taskdetail, b.AfterBakChan)
		if err != nil {
			log.Logger.Error(err)
			return
		}
		log.Logger.Infof("StartBakAll开启备份任务TASK:%s", task.DBName)
		if err = bakhandler.StartBak(); err != nil {
			log.Logger.Error(err)
			return
		}
		// 修改任务启动状态
		taskinfo.Status = 1
		taskinfo.Id = task.Id
		if err = taskinfo.UpdatesStatus(database.GetDB()); err != nil {
			log.Logger.Error(err)
			return
		}
	}
	middleware.ResponseSuccess(c, "批量启动任务成功")
}

func (bak *BakController) StartBakAllByHost(c *gin.Context) {
	params := &dto.HostIDInput{}
	if err := params.BindValidParams(c); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	var b BakController
	taskinfo := &dao.TaskInfo{}
	result, err := taskinfo.FindAllTask(c, database.GetDB(), params)
	if err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 2002, err)
		return
	}
	if len(result) == 0 {
		log.Logger.Info("当前主机备份任务为空")
		middleware.ResponseError(c, 2003, errors.New("当前主机备份任务为空"))
		return
	}
	b.AfterBakChan = make(chan *core.BakHandler, 10)
	go b.ListenAndSave(c, database.GetDB(), b.AfterBakChan)
	for _, task := range result {
		if task.Status == 1 {
			log.Logger.Infof("TASK_ID:%d,HOSTID:%d,%s数据库任务已经打开,返回", task.Id, task.HostID, task.DBName)
			continue
		}

		taskdetail, err := taskinfo.TaskDetail(c, database.GetDB(), task)
		if err != nil {
			log.Logger.Error(err)
			return
		}
		bakhandler, err := core.NewBakHandler(taskdetail, b.AfterBakChan)
		if err != nil {
			log.Logger.Error(err)
			return
		}
		if err = bakhandler.StartBak(); err != nil {
			log.Logger.Error(err)
			return
		}
		// 修改任务启动状态
		taskinfo.Status = 1
		taskinfo.Id = task.Id
		if err = taskinfo.UpdatesStatus(database.GetDB()); err != nil {
			log.Logger.Error(err)
			return
		}
	}
	middleware.ResponseSuccess(c, "批量启动任务成功")
}

func (bak *BakController) StopBak(c *gin.Context) {
	params := &dto.Bak{}
	if err := params.BindValidParams(c); err != nil {
		log.Logger.Error("BakHandleController 解析参数失败")
		middleware.ResponseError(c, 1007, err)
		return
	}
	var err error
	//将任务状态改为false
	taskinfo := &dao.TaskInfo{
		Id: params.ID,
	}
	taskinfo, err = taskinfo.Find(c, database.GetDB(), taskinfo)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	taskinfo.Status = 0
	if err = taskinfo.UpdatesStatus(database.GetDB()); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	var bakhandler = &core.BakHandler{}
	if err = bakhandler.StopBak(params.ID); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 1008, err)
		return
	}
	middleware.ResponseSuccess(c, "任务停止成功")
}

func (bak *BakController) StopBakAll(c *gin.Context) {
	taskinfo := &dao.TaskInfo{}
	result, err := taskinfo.FindAllTask(c, database.GetDB(), nil)
	if err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 2002, err)
		return
	}
	if len(result) == 0 {
		log.Logger.Info("当前主机备份任务为空")
		middleware.ResponseError(c, 2003, errors.New("当前主机备份任务为空"))
		return
	}
	for _, task := range result {
		if task.Status == 0 {
			log.Logger.Infof("TASK_ID:%d,HOSTID:%d,%s数据库任务已经关闭,返回", task.Id, task.HostID, task.DBName)
			continue
		}
		bakhandler := &core.BakHandler{DbName: task.DBName}
		err = bakhandler.StopBak(task.Id)
		if err != nil {
			log.Logger.Warning(err)
			middleware.ResponseError(c, 2004, err)
			return
		}
		// 修改任务启动状态为关闭
		taskinfo.Id = task.Id
		taskinfo.Status = 0
		if err = taskinfo.UpdatesStatus(database.GetDB()); err != nil {
			log.Logger.Error(err)
			middleware.ResponseError(c, 2005, err)
			return
		}
	}
	middleware.ResponseSuccess(c, "批量停止任务成功")
}

func (bak *BakController) StopBakAllByHost(c *gin.Context) {
	params := &dto.HostIDInput{}
	if err := params.BindValidParams(c); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	taskinfo := &dao.TaskInfo{}
	result, err := taskinfo.FindAllTask(c, database.GetDB(), params)
	if err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 2002, err)
		return
	}
	if len(result) == 0 {
		log.Logger.Info("当前主机备份任务为空")
		middleware.ResponseError(c, 2003, errors.New("当前主机备份任务为空"))
		return
	}
	for _, task := range result {
		if task.Status == 0 {
			log.Logger.Infof("TASK_ID:%d,HOSTID:%d,%s数据库任务已经关闭,返回", task.Id, task.HostID, task.DBName)
			continue
		}
		bakhandler := &core.BakHandler{DbName: task.DBName}
		err = bakhandler.StopBak(task.Id)
		if err != nil {
			log.Logger.Warning(err)
			middleware.ResponseError(c, 2004, err)
			return
		}
		// 修改任务启动状态为关闭
		taskinfo.Id = task.Id
		taskinfo.Status = 0
		if err = taskinfo.UpdatesStatus(database.GetDB()); err != nil {
			log.Logger.Error(err)
			middleware.ResponseError(c, 2005, err)
			return
		}
	}
	middleware.ResponseSuccess(c, "批量停止任务成功")
}

func (bak *BakController) ListenAndSave(ctx *gin.Context, tx *gorm.DB, AfterBakChan chan *core.BakHandler) {
	log.Logger.Info("开始监听备份状态消息")
	for {
		select {
		case afterbakhandler := <-AfterBakChan:
			bakhistory := &dao.BakHistory{
				TaskID:     afterbakhandler.TaskID,
				Host:       afterbakhandler.Host,
				DBName:     afterbakhandler.DbName,
				BakStatus:  afterbakhandler.BakStatus,
				Msg:        afterbakhandler.BakMsg,
				OssStatus:  afterbakhandler.OssStatus,
				DingStatus: afterbakhandler.DingStatus,
				FileSize:   afterbakhandler.FileSize,
				FileName:   afterbakhandler.FileName,
				BakTime:    time.Now(),
			}
			log.Logger.Info("接收到备份消息，数据入库")
			if err := bakhistory.Save(ctx, tx); err != nil {
				tx.Rollback()
				log.Logger.Error("保存备份历史到数据库失败", err)
			}
		}
	}
}
