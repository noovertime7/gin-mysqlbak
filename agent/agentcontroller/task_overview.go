package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/agentservice"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

func AgentOverViewTaskRegister(group *gin.RouterGroup) {
	task := &OverViewTask{service: agentservice.GetClusterTaskOverViewService()}
	group.GET("/overview_task_list", task.GetOverViewTaskList)
	group.GET("/overview_task_start", task.StartOveTask)
	group.GET("/overview_task_stop", task.StopOveTask)
	group.DELETE("/overview_task_delete", task.DeleteTask)
	group.PUT("/overview_task_restore", task.RestoreTask)
	group.GET("/overview_task_sync", task.Sync)
}

type OverViewTask struct {
	service *agentservice.TaskOverViewService
}

func (o *OverViewTask) GetOverViewTaskList(ctx *gin.Context) {
	params := &agentdto.TaskOverViewListInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := o.service.GetTaskOverViewList(ctx, params)
	if err != nil {
		log.Logger.Error("agent获取任务总览", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskOverViewGetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (o *OverViewTask) StartOveTask(ctx *gin.Context) {
	params := &agentdto.StartOverViewBakInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := o.service.StartTask(ctx, params)
	if err != nil {
		log.Logger.Error("agent启动任务成功", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.BakStartError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (o *OverViewTask) StopOveTask(ctx *gin.Context) {
	params := &agentdto.StopOverViewBakInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := o.service.StopTask(ctx, params)
	if err != nil {
		log.Logger.Error("agent启动任务失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.BakStopError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (o *OverViewTask) DeleteTask(ctx *gin.Context) {
	params := &agentdto.DeleteOverViewTaskInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := o.service.DeleteTask(ctx, params)
	if err != nil {
		log.Logger.Error("agent删除任务失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.BakStopError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (o *OverViewTask) RestoreTask(ctx *gin.Context) {
	params := &agentdto.DeleteOverViewTaskInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := o.service.RestoreTask(ctx, params)
	if err != nil {
		log.Logger.Error("agent还原任务失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskRestoreError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (o *OverViewTask) Sync(ctx *gin.Context) {
	if err := o.service.Store(ctx); err != nil {
		log.Logger.Error("手动同步成功")
	}
	log.Logger.Info("手动同步成功")
	middleware.ResponseSuccess(ctx, "操作成功")
}
