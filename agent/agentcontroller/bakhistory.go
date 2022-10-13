package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/agentservice"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type BakHistoryController struct {
	historyService *agentservice.HistoryService
}

func BakHistoryRegister(group *gin.RouterGroup) {
	history := &BakHistoryController{historyService: agentservice.GetClusterHistoryService()}
	group.GET("/historylist", history.HistoryList)
	group.GET("/history_num_info", history.GetHistoryNumInfo)
	group.DELETE("/historydelete", history.DeleteHistory)
}

func (b *BakHistoryController) HistoryList(ctx *gin.Context) {
	params := &agentdto.HistoryListInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := b.historyService.GetHistoryList(ctx, params)
	if err != nil {
		log.Logger.Error("agent获取历史记录列表失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (b *BakHistoryController) GetHistoryNumInfo(ctx *gin.Context) {
	params := &agentdto.HistoryServiceNameInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := b.historyService.GetHistoryNumInfo(ctx, params)
	if err != nil {
		log.Logger.Error("获取历史记录数量信息失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.HistoryGetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (b *BakHistoryController) DeleteHistory(ctx *gin.Context) {
	params := &agentdto.BakHistoryDeleteInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := b.historyService.DeleteHistory(ctx, params)
	if err != nil || !data.OK {
		log.Logger.Error("删除失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data.Message)
}
