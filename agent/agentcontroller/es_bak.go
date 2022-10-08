package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/agentservice"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type EsBakController struct {
	esBakService *agentservice.EsBakService
}

func EsBakRegister(group *gin.RouterGroup) {
	esBak := &EsBakController{
		esBakService: agentservice.GetClusterEsBakService(),
	}
	group.PUT("/start", esBak.Start)
	group.PUT("/stop", esBak.Stop)
}

func (e *EsBakController) Start(ctx *gin.Context) {
	params := &agentdto.ESBakStartInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := e.esBakService.StartEsBak(ctx, params)
	if err != nil || !data.OK {
		log.Logger.Error("es启动任务失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.BakStartError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data.Message)
	log.Logger.Info("es启动任务成功", data.Message)
}

func (e *EsBakController) Stop(ctx *gin.Context) {
	params := &agentdto.ESBakStopInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := e.esBakService.StopEsBak(ctx, params)
	if err != nil || !data.OK {
		log.Logger.Error("es停止任务失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.BakStopError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data.Message)
	log.Logger.Info("es停止任务成功", data.Message)
}
