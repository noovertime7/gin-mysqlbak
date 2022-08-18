package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/bak"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type BakController struct{}

func BakRegister(group *gin.RouterGroup) {
	bak := &BakController{}
	group.PUT("/bakstart", bak.StartBak)
	group.PUT("/bakstop", bak.StopBak)
	group.PUT("/bakhoststart", bak.StartBakByHost)
	group.PUT("/bakhoststop", bak.StopBakByHost)
}

func (b *BakController) StartBak(ctx *gin.Context) {
	params := &agentdto.StartBakInput{}
	if err := params.BindValidParm(ctx); err != nil {
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	bakService := pkg.GetBakService(params.ServiceName).(bak.BakService)
	bakStartInput := &bak.StartBakInput{
		TaskID:      params.TaskID,
		ServiceName: params.ServiceName,
	}
	log.Logger.Info("agent开始启动任务", bakStartInput)
	data, err := bakService.StartBak(ctx, bakStartInput)
	if err != nil || !data.OK {
		middleware.ResponseError(ctx, 30001, err)
		return
	}
	middleware.ResponseSuccess(ctx, "任务启动成功")
}
func (b *BakController) StopBak(ctx *gin.Context) {
	params := &agentdto.StopBakInput{}
	if err := params.BindValidParm(ctx); err != nil {
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	bakService := pkg.GetBakService(params.ServiceName).(bak.BakService)
	bakStartInput := &bak.StopBakInput{
		TaskID:      params.TaskID,
		ServiceName: params.ServiceName,
	}
	log.Logger.Info("agent开始停止任务", bakStartInput)
	data, err := bakService.StopBak(ctx, bakStartInput)
	if err != nil || !data.OK {
		middleware.ResponseError(ctx, 30001, err)
		return
	}
	middleware.ResponseSuccess(ctx, "任务停止成功")
}
func (b *BakController) StartBakByHost(ctx *gin.Context) {
	params := &agentdto.StartBakByHostInput{}
	if err := params.BindValidParm(ctx); err != nil {
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	bakService := pkg.GetBakService(params.ServiceName).(bak.BakService)
	bakStartInput := &bak.StartBakByHostInput{
		HostID:      params.HostID,
		ServiceName: params.ServiceName,
	}
	log.Logger.Info("agent开始启动所有Host任务", bakStartInput)
	data, err := bakService.StartBakByHost(ctx, bakStartInput)
	if err != nil || !data.OK {
		middleware.ResponseError(ctx, 30001, err)
		return
	}
	middleware.ResponseSuccess(ctx, "任务启动成功")
}
func (b *BakController) StopBakByHost(ctx *gin.Context) {
	params := &agentdto.StopBakByHostInput{}
	if err := params.BindValidParm(ctx); err != nil {
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	bakService := pkg.GetBakService(params.ServiceName).(bak.BakService)
	bakStartInput := &bak.StopBakByHostInput{
		HostID:      params.HostID,
		ServiceName: params.ServiceName,
	}
	log.Logger.Info("agent开始停止所有Host任务", bakStartInput)
	data, err := bakService.StopBakByHost(ctx, bakStartInput)
	if err != nil || !data.OK {
		middleware.ResponseError(ctx, 30001, err)
		return
	}
	middleware.ResponseSuccess(ctx, "任务停止成功")
}