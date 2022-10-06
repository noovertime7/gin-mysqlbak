package agentservice

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/host"
	"sync"
)

var (
	clusterHostService *HostService
	HostServiceOnce    sync.Once
)

// GetClusterHostService  单例模式
func GetClusterHostService() *HostService {
	HostServiceOnce.Do(func() {
		clusterHostService = &HostService{}
	})
	return clusterHostService
}

type HostService struct{}

func (h *HostService) AddHost(c *gin.Context, params *agentdto.HostAddInput) (*host.HostOneMessage, error) {
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	//如果类型为es添加http
	if params.Type == 2 {
		params.Host = "http://" + params.Host
	}
	return hostService.AddHost(c, &host.HostAddInput{
		Host:     params.Host,
		UserName: params.User,
		Password: params.Password,
		Content:  params.Content,
		Type:     params.Type,
	}, ops)
}

func (h *HostService) DeleteHost(c *gin.Context, params *agentdto.HostDeleteInput) (*host.HostOneMessage, error) {
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	return hostService.DeleteHost(c, &host.HostIDInput{
		ID: int64(params.ID),
	}, ops)
}

func (h *HostService) UpdateHost(c *gin.Context, params *agentdto.HostUpdateInput) (*host.HostOneMessage, error) {
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	//如果类型为es添加http
	if params.Type == 2 {
		params.Host = "http://" + params.Host
	}
	return hostService.UpdateHost(c, &host.HostUpdateInput{
		ID:       int64(params.ID),
		Host:     params.Host,
		UserName: params.User,
		Password: params.Password,
		Content:  params.Content,
		Type:     params.Type,
	}, ops)
}

func (h *HostService) HostList(c *gin.Context, params *agentdto.HostListInput) (*host.HostListOutPut, error) {
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	return hostService.GetHostList(c, &host.HostListInput{
		Info:     params.Info,
		PageNo:   params.PageNo,
		PageSize: params.PageSize,
	}, ops)
}

func (h *HostService) HostTest(c *gin.Context, params *agentdto.HostIDInput) (*host.HostOneMessage, error) {
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	return hostService.TestHost(c, &host.HostIDInput{ID: params.HostID}, ops)
}

func (h *HostService) GetHostNames(ctx *gin.Context, params *agentdto.HostNamesInput) (*host.HostNames, error) {
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	return hostService.GetHostNames(ctx, &host.HostNamesInput{Type: params.Type}, ops)
}
