package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/host"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type AgentHostController struct{}

func AgentHostRegister(group *gin.RouterGroup) {
	agenthost := &AgentHostController{}
	group.GET("/hostlist", agenthost.HostList)
	group.POST("/hostadd", agenthost.AddHost)
	group.DELETE("/hostdelete", agenthost.DeleteHost)
	group.PUT("/hostupdate", agenthost.UpdateHost)
}

func (a *AgentHostController) AddHost(c *gin.Context) {
	params := &agentdto.HostAddInput{}
	if err := params.BindValidParams(c); err != nil {
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(c, 30001, err)
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	data, err := hostService.AddHost(c, &host.HostAddInput{
		Host:     params.Host,
		UserName: params.User,
		Password: params.Password,
		Content:  params.Content,
	}, ops)
	if err != nil || !data.OK {
		log.Logger.Error("agent添加主机失败", err)
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(c, data.Message)
	log.Logger.Info("agent添加主机成功")
}

func (a *AgentHostController) DeleteHost(c *gin.Context) {
	params := &agentdto.HostDeleteInput{}
	if err := params.BindValidParams(c); err != nil {
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(c, 30001, err)
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	data, err := hostService.DeleteHost(c, &host.HostDeleteInput{
		ID: int64(params.ID),
	}, ops)
	if err != nil || !data.OK {
		log.Logger.Error("agent删除主机失败", err)
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(c, data.Message)
	log.Logger.Info("agent删除主机成功")
}

func (a *AgentHostController) UpdateHost(c *gin.Context) {
	params := &agentdto.HostUpdateInput{}
	if err := params.BindValidParams(c); err != nil {
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(c, 30001, err)
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	data, err := hostService.UpdateHost(c, &host.HostUpdateInput{
		ID:       int64(params.ID),
		Host:     params.Host,
		UserName: params.User,
		Password: params.Password,
		Content:  params.Content,
	}, ops)
	if err != nil || !data.OK {
		log.Logger.Error("agent更新主机失败", err)
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(c, data.Message)
	log.Logger.Info("agent更新主机成功")
}

func (a *AgentHostController) HostList(c *gin.Context) {
	params := &agentdto.HostListInput{}
	if err := params.BindValidParams(c); err != nil {
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(c, 30001, err)
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	out, err := hostService.GetHostList(c, &host.HostListInput{
		Info:     params.Info,
		PageNo:   int64(params.PageNo),
		PageSize: int64(params.PageSize),
	}, ops)
	if err != nil {
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(c, out)
}
