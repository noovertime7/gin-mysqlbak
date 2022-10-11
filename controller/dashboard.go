package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
	"github.com/noovertime7/gin-mysqlbak/services/local"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type DashboardController struct {
	dashBoardService *local.DashboardService
}

func DashboardRegister(group *gin.RouterGroup) {
	dashboard := &DashboardController{dashBoardService: local.GetDashboardService()}
	group.GET("/service_task_num", dashboard.GetSvcTNum)
	group.GET("/service_info_by_date", dashboard.GetInfoByDate)
}

func (d *DashboardController) GetSvcTNum(ctx *gin.Context) {
	data, err := d.dashBoardService.GetSvcTNum(ctx)
	if err != nil {
		log.Logger.Error("获取dashboard数据失败")
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ServerError, err))
	}
	middleware.ResponseSuccess(ctx, data)
}

func (d *DashboardController) GetInfoByDate(ctx *gin.Context) {
	params := &dto.AgentDateInfoInput{}
	if err := params.BindValidParam(ctx); err != nil {
		log.Logger.Error("绑定失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := d.dashBoardService.GetTaskNumByDate(ctx, params)
	if err != nil {
		log.Logger.Error("获取dashboard数据失败")
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ServerError, err))
	}
	middleware.ResponseSuccess(ctx, data)
}
