package agentservice

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/bakhistory"
	"sync"
)

var (
	clusterHistoryService *HistoryService
	HistoryServiceOnce    sync.Once
)

// GetClusterHistoryService  单例模式
func GetClusterHistoryService() *HistoryService {
	HistoryServiceOnce.Do(func() {
		clusterHistoryService = &HistoryService{}
	})
	return clusterHistoryService
}

type HistoryService struct {
}

func (h *HistoryService) GetHistoryList(ctx *gin.Context, params *agentdto.HistoryListInput) (*bakhistory.HistoryListOutput, error) {
	historyService, addr, err := pkg.GetHistoryService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	historyListInput := &bakhistory.HistoryListInput{
		Info:      params.Info,
		PageNo:    params.PageNo,
		PageSize:  params.PageSize,
		SortField: params.SortField,
		SortOrder: params.SortOrder,
		Status:    params.Status,
	}
	return historyService.GetHistoryList(ctx, historyListInput, ops)
}

func (h *HistoryService) GetHistoryNumInfo(ctx *gin.Context, params *agentdto.HistoryServiceNameInput) (*bakhistory.HistoryNumInfoOut, error) {
	historyService, addr, err := pkg.GetHistoryService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	return historyService.GetHistoryNumInfo(ctx, &bakhistory.Empty{}, ops)
}

func (h *HistoryService) DeleteHistory(ctx *gin.Context, params *agentdto.BakHistoryDeleteInput) (*bakhistory.HistoryOneMessage, error) {
	historyService, addr, err := pkg.GetHistoryService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	return historyService.DeleteHistory(ctx, &bakhistory.HistoryIDInput{ID: params.ID}, ops)
}
