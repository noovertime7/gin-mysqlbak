package agentservice

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/esbak"
	"sync"
)

var (
	clusterEsBakService *EsBakService
	EsBakServiceOnce    sync.Once
)

// GetClusterEsBakService  单例模式
func GetClusterEsBakService() *EsBakService {
	EsBakServiceOnce.Do(func() {
		clusterEsBakService = &EsBakService{}
	})
	return clusterEsBakService
}

type EsBakService struct{}

func (e *EsBakService) StartEsBak(ctx *gin.Context, params *agentdto.ESBakStartInput) (*esbak.EsBakOneMessage, error) {
	EsBakService, addr, err := pkg.GetESBakService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	return EsBakService.Start(ctx, &esbak.StartEsBakInput{TaskID: params.ID}, ops)
}

func (e *EsBakService) StopEsBak(ctx *gin.Context, params *agentdto.ESBakStopInput) (*esbak.EsBakOneMessage, error) {
	EsBakService, addr, err := pkg.GetESBakService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	return EsBakService.Stop(ctx, &esbak.StopEsBakInput{TaskID: params.ID}, ops)
}
