package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/core"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/public/database"
	"github.com/noovertime7/gin-mysqlbak/services/local"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/pkg/errors"
)

type BakController struct{}

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
	params := &dto.Bak{}
	if err := params.BindValidParams(c); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	service := local.GetBakBakService()
	if err := service.Start(c); err != nil {
		log.Logger.Error("启动任务失败", err)
		middleware.ResponseError(c, 10002, err)
		return
	}
	middleware.ResponseSuccess(c, "启动任务成功")
}

func (bak *BakController) StartBakAll(c *gin.Context) {
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
	tx := database.GetDB()
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
		bakhandler, err := core.NewBakHandler(taskdetail)
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
	service := local.GetBakBakService()
	if err := service.StartTaskByHost(c, params); err != nil {
		log.Logger.Error("启动主机下所有任务失败")
		middleware.ResponseError(c, 20001, err)
		return
	}
	middleware.ResponseSuccess(c, "批量启动任务成功")
}

func (bak *BakController) StopBak(c *gin.Context) {
	params := &dto.Bak{}
	if err := params.BindValidParams(c); err != nil {
		log.Logger.Error("BakHandleController 解析参数失败")
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	service := local.GetBakBakService()
	if err := service.Stop(c, params); err != nil {
		log.Logger.Error("停止失败", err)
		middleware.ResponseError(c, 10002, err)
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
	service := local.GetBakBakService()
	if err := service.StopTaskByHost(c, params); err != nil {
		log.Logger.Error("根据主机停止任务失败")
		middleware.ResponseError(c, 20003, err)
		return
	}
	middleware.ResponseSuccess(c, "批量停止任务成功")
}
