package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/agentservice"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type EsTaskController struct {
	service *agentservice.EsTaskService
}

func EsTaskRegister(group *gin.RouterGroup) {
	esTask := &EsTaskController{service: agentservice.GetClusterEsTaskService()}
	group.POST("/taskadd", esTask.AddEsTask)
	group.DELETE("/taskdelete", esTask.DeleteEsTask)
	group.PUT("/taskupdate", esTask.UpdateEsTask)
	group.GET("/tasklist", esTask.GetEsTaskList)
	group.GET("/taskdetail", esTask.GetEsTaskDetail)
	group.GET("/task_restore", esTask.RestoreEsTask)
}

func (e *EsTaskController) AddEsTask(ctx *gin.Context) {
	params := &agentdto.ESBakTaskAddInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := e.service.AddEsTask(ctx, params)
	if err != nil || !data.OK {
		log.Logger.Error("es添加任务失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskAddError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data.Message)
	log.Logger.Info("es添加任务成功", data.Message)
}

func (e *EsTaskController) DeleteEsTask(ctx *gin.Context) {
	params := &agentdto.ESBakTaskIDInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := e.service.DeleteEsTask(ctx, params)
	if err != nil || !data.OK {
		log.Logger.Error("es删除任务失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskDeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data.Message)
	log.Logger.Info("es删除任务成功", data.Message)
}

func (e *EsTaskController) RestoreEsTask(ctx *gin.Context) {
	params := &agentdto.ESBakTaskIDInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := e.service.RestoreEsTask(ctx, params)
	if err != nil || !data.OK {
		log.Logger.Error("es删除任务失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskDeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data.Message)
	log.Logger.Info("es删除任务成功", data.Message)
}

func (e *EsTaskController) UpdateEsTask(ctx *gin.Context) {
	params := &agentdto.ESBakTaskUpdateInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := e.service.UpdateEsTask(ctx, params)
	if err != nil || !data.OK {
		log.Logger.Error("es修改任务失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskUpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data.Message)
	log.Logger.Info("es修改任务成功", data.Message)
}

func (e *EsTaskController) GetEsTaskList(ctx *gin.Context) {
	params := &agentdto.ESBakTaskListInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := e.service.GetEsTaskList(ctx, params)
	if err != nil {
		log.Logger.Error("获取es_task列表失败")
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskGetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
func (e *EsTaskController) GetEsTaskDetail(ctx *gin.Context) {
	params := &agentdto.ESBakTaskIDInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := e.service.GetEsTaskDetail(ctx, params)
	if err != nil {
		log.Logger.Error("获取Agent详情失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskGetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
