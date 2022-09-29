package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/esbak"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type EsBakController struct{}

func EsBakRegister(group *gin.RouterGroup) {
	esBak := &EsBakController{}
	group.PUT("/start", esBak.Start)
	group.PUT("/stop", esBak.Stop)
}

func (e *EsBakController) Start(ctx *gin.Context) {
	params := &agentdto.ESBakStartInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	EsBakService, addr, err := pkg.GetESBakService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.AgentGetAddressError, err))
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	data, err := EsBakService.Start(ctx, &esbak.StartEsBakInput{TaskID: params.ID}, ops)
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
	EsBakService, addr, err := pkg.GetESBakService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.AgentGetAddressError, err))
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	data, err := EsBakService.Stop(ctx, &esbak.StopEsBakInput{TaskID: params.ID}, ops)
	if err != nil || !data.OK {
		log.Logger.Error("es停止任务失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.BakStopError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data.Message)
	log.Logger.Info("es停止任务成功", data.Message)
}
