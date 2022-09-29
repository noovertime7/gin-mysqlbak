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

type EsHistoryController struct{}

func EsHistoryRegister(group *gin.RouterGroup) {
	esHistory := &EsHistoryController{}
	group.DELETE("/historydelete", esHistory.DeleteEsHistory)
	group.GET("/historylist", esHistory.GetEsHistoryList)
}

func (e *EsHistoryController) DeleteEsHistory(ctx *gin.Context) {
	params := &agentdto.EsHistoryIDInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	EsHistoryService, addr, err := pkg.GetESHistoryService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.AgentGetAddressError, err))
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	data, err := EsHistoryService.DeleteESHistory(ctx, &esbak.ESHistoryIDInput{ID: params.ID}, ops)
	if err != nil || !data.OK {
		log.Logger.Error("es删除历史记录失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.HistoryDeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data.Message)
	log.Logger.Info("es删除历史记录成功", data.Message)
}

func (e *EsHistoryController) GetEsHistoryList(ctx *gin.Context) {
	params := &agentdto.ESHistoryListInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	EsHistoryService, addr, err := pkg.GetESHistoryService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.AgentGetAddressError, err))
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	data, err := EsHistoryService.GetEsHistoryList(ctx, &esbak.GetEsHistoryListInput{
		Info:      params.Info,
		PageNo:    params.PageNo,
		PageSize:  params.PageSize,
		SortOrder: params.SortOrder,
		SortField: params.SortField,
		Status:    params.Status,
	}, ops)
	if err != nil {
		log.Logger.Error("获取es_task历史记录失败")
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.HistoryGetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
