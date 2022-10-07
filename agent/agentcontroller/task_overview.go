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
