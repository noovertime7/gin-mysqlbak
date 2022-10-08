package agentservice

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/bak"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"sync"
)

var (
	clusterBakService *BakService
	BakServiceOnce    sync.Once
)

// GetClusterBakService  单例模式
func GetClusterBakService() *BakService {
	BakServiceOnce.Do(func() {
		clusterBakService = &BakService{}
	})
	return clusterBakService
}

type BakService struct{}

func (b *BakService) StartBak(ctx *gin.Context, params *agentdto.StartBakInput) (*bak.BakOneMessage, error) {
	bakService, addr, err := pkg.GetBakService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	bakStartInput := &bak.StartBakInput{
		TaskID:      params.TaskID,
		ServiceName: params.ServiceName,
	}
	log.Logger.Info("agent开始启动任务", bakStartInput)
	return bakService.StartBak(ctx, bakStartInput, ops)
}

func (b *BakService) StopBak(ctx *gin.Context, params *agentdto.StopBakInput) (*bak.BakOneMessage, error) {
	bakService, addr, err := pkg.GetBakService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	bakStartInput := &bak.StopBakInput{
		TaskID:      params.TaskID,
		ServiceName: params.ServiceName,
	}
	log.Logger.Info("agent开始停止任务", bakStartInput, addr)
	return bakService.StopBak(ctx, bakStartInput, ops)
}

func (b *BakService) StartBakByHost(ctx *gin.Context, params *agentdto.StartBakByHostInput) (*bak.BakOneMessage, error) {
	bakService, addr, err := pkg.GetBakService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	bakStartInput := &bak.StartBakByHostInput{
		HostID:      params.HostID,
		ServiceName: params.ServiceName,
	}
	log.Logger.Info("agent开始启动所有Host任务", bakStartInput)
	return bakService.StartBakByHost(ctx, bakStartInput, ops)
}
func (b *BakService) StopBakByHost(ctx *gin.Context, params *agentdto.StopBakByHostInput) (*bak.BakOneMessage, error) {
	bakService, addr, err := pkg.GetBakService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	bakStartInput := &bak.StopBakByHostInput{
		HostID:      params.HostID,
		ServiceName: params.ServiceName,
	}
	log.Logger.Info("agent开始停止所有Host任务", bakStartInput)
	return bakService.StopBakByHost(ctx, bakStartInput, ops)
}
