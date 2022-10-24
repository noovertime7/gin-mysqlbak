package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/agentservice"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
)

type BakController struct {
	bakService *agentservice.BakService
}

func BakRegister(group *gin.RouterGroup) {
	bak := &BakController{
		bakService: agentservice.GetClusterBakService(),
	}
	group.PUT("/bakstart", bak.StartBak)
	group.PUT("/bak_test", bak.TestBak)
	group.PUT("/bakstop", bak.StopBak)
	group.PUT("/bakhoststart", bak.StartBakByHost)
	group.PUT("/bakhoststop", bak.StopBakByHost)
}

func (b *BakController) StartBak(ctx *gin.Context) {
	params := &agentdto.StartBakInput{}
	if err := params.BindValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := b.bakService.StartBak(ctx, params)
	if err != nil || !data.OK {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.BakStartError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "任务启动成功")
}

func (b *BakController) TestBak(ctx *gin.Context) {
	params := &agentdto.StartBakInput{}
	if err := params.BindValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := b.bakService.TestBak(ctx, params)
	if err != nil || !data.OK {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.BakStartError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "测试任务启动成功")
}

func (b *BakController) StopBak(ctx *gin.Context) {
	params := &agentdto.StopBakInput{}
	if err := params.BindValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := b.bakService.StopBak(ctx, params)
	if err != nil || !data.OK {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.BakStopError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "任务停止成功")
}
func (b *BakController) StartBakByHost(ctx *gin.Context) {
	params := &agentdto.StartBakByHostInput{}
	if err := params.BindValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := b.bakService.StartBakByHost(ctx, params)
	if err != nil || !data.OK {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.BakStartAllError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "任务启动成功")
}
func (b *BakController) StopBakByHost(ctx *gin.Context) {
	params := &agentdto.StopBakByHostInput{}
	if err := params.BindValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := b.bakService.StopBakByHost(ctx, params)
	if err != nil || !data.OK {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.BakStopAllError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "任务停止成功")
}
