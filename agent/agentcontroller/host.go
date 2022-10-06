package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/agentservice"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type AgentHostController struct {
	service *agentservice.HostService
}

func AgentHostRegister(group *gin.RouterGroup) {
	agenthost := &AgentHostController{service: agentservice.GetClusterHostService()}
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
	data, err := a.service.AddHost(c, params)
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
	data, err := a.service.DeleteHost(c, params)
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
	data, err := a.service.UpdateHost(c, params)
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
	out, err := a.service.HostList(c, params)
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
	data, err := a.service.HostTest(c, params)
	if err != nil {
		log.Logger.Error("主机测试失败")
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.ServerError, err))
	}
	middleware.ResponseSuccess(c, data.Message)
}

func (a *AgentHostController) GetHostNames(ctx *gin.Context) {
	params := &agentdto.HostNamesInput{}
	if err := params.BindValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := a.service.GetHostNames(ctx, params)
	if err != nil {
		log.Logger.Error("获取主机名称列表失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.HostGetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
