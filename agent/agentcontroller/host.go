package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/host"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type AgentHostController struct{}

func AgentHostRegister(group *gin.RouterGroup) {
	agenthost := &AgentHostController{}
	group.GET("/hostlist", agenthost.HostList)
	group.GET("/host_test", agenthost.HostTest)
	group.GET("/host_names", agenthost.GetHostNames)
	group.POST("/hostadd", agenthost.AddHost)
	group.DELETE("/hostdelete", agenthost.DeleteHost)
	group.PUT("/hostupdate", agenthost.UpdateHost)
}

func (a *AgentHostController) AddHost(c *gin.Context) {
	params := &agentdto.HostAddInput{}
	if err := params.BindValidParams(c); err != nil {
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.AgentGetAddressError, err))
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	//如果类型为es添加http
	if params.Type == 2 {
		params.Host = "http://" + params.Host
	}
	data, err := hostService.AddHost(c, &host.HostAddInput{
		Host:     params.Host,
		UserName: params.User,
		Password: params.Password,
		Content:  params.Content,
		Type:     params.Type,
	}, ops)
	if err != nil || !data.OK {
		log.Logger.Error("agent添加主机失败", err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.HostAddError, err))
		return
	}
	middleware.ResponseSuccess(c, data.Message)
	log.Logger.Info("agent添加主机成功")
}

func (a *AgentHostController) DeleteHost(c *gin.Context) {
	params := &agentdto.HostDeleteInput{}
	if err := params.BindValidParams(c); err != nil {
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.AgentGetAddressError, err))
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	data, err := hostService.DeleteHost(c, &host.HostIDInput{
		ID: int64(params.ID),
	}, ops)
	if err != nil || !data.OK {
		log.Logger.Error("agent删除主机失败", err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.HostDeleteError, err))
		return
	}
	middleware.ResponseSuccess(c, data.Message)
	log.Logger.Info("agent删除主机成功")
}

func (a *AgentHostController) UpdateHost(c *gin.Context) {
	params := &agentdto.HostUpdateInput{}
	if err := params.BindValidParams(c); err != nil {
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.AgentGetAddressError, err))
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	//如果类型为es添加http
	if params.Type == 2 {
		params.Host = "http://" + params.Host
	}
	data, err := hostService.UpdateHost(c, &host.HostUpdateInput{
		ID:       int64(params.ID),
		Host:     params.Host,
		UserName: params.User,
		Password: params.Password,
		Content:  params.Content,
		Type:     params.Type,
	}, ops)
	if err != nil || !data.OK {
		log.Logger.Error("agent更新主机失败", err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.HostUpdateError, err))
		return
	}
	middleware.ResponseSuccess(c, data.Message)
	log.Logger.Info("agent更新主机成功")
}

func (a *AgentHostController) HostList(c *gin.Context) {
	params := &agentdto.HostListInput{}
	if err := params.BindValidParams(c); err != nil {
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.AgentGetAddressError, err))
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	out, err := hostService.GetHostList(c, &host.HostListInput{
		Info:     params.Info,
		PageNo:   params.PageNo,
		PageSize: params.PageSize,
	}, ops)
	if err != nil {
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.HostGetError, err))
		return
	}
	middleware.ResponseSuccess(c, out)
}

func (a *AgentHostController) HostTest(c *gin.Context) {
	params := &agentdto.HostIDInput{}
	if err := params.BindValidParams(c); err != nil {
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.AgentGetAddressError, err))
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	data, err := hostService.TestHost(c, &host.HostIDInput{ID: params.HostID}, ops)
	if err != nil || !data.OK {
		log.Logger.Error("agent添加主机失败", err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.ServerError, err))
		return
	}
	middleware.ResponseSuccess(c, data.Message)
}

func (a *AgentHostController) GetHostNames(ctx *gin.Context) {
	params := &agentdto.HostNamesInput{}
	if err := params.BindValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	hostService, addr, err := pkg.GetHostService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.AgentGetAddressError, err))
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	data, err := hostService.GetHostNames(ctx, &host.HostNamesInput{Type: params.Type}, ops)
	if err != nil {
		log.Logger.Error("获取主机名称列表失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.HostGetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
