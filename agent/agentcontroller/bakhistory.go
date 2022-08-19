package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/bakhistory"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type BakHistoryController struct{}

func BakHistoryRegister(group *gin.RouterGroup) {
	history := &BakHistoryController{}
	group.GET("/historylist", history.HistoryList)
	group.DELETE("/historydelete", history.DeleteHistory)
}

func (b *BakHistoryController) HistoryList(ctx *gin.Context) {
	params := &agentdto.HistoryListInput{}
	if err := params.BindValidParm(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}

	historyService, addr, err := pkg.GetHistoryService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	historyListInput := &bakhistory.HistoryListInput{
		Info:     params.Info,
		PageNo:   params.PageNo,
		PageSize: params.PageSize,
		Sort:     params.Sort,
	}
	data, err := historyService.GetHistoryList(ctx, historyListInput, ops)
	if err != nil {
		log.Logger.Error("agent获取历史记录列表失败", err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (b *BakHistoryController) DeleteHistory(ctx *gin.Context) {
	params := &agentdto.BakHistoryDeleteInput{}
	if err := params.BindValidParm(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	historyService, addr, err := pkg.GetHistoryService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, 30001, err)
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}

	data, err := historyService.DeleteHistory(ctx, &bakhistory.HistoryIDInput{ID: params.ID}, ops)
	if err != nil || !data.OK {
		log.Logger.Error("删除失败", err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(ctx, data.Message)
}
