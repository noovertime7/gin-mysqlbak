package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/agentservice"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type AgentController struct {
}

var AgentService *agentservice.AgentService

func AgentRegister(group *gin.RouterGroup) {
	agent := &AgentController{}
	group.GET("/agentlist", agent.GetAgentList)
	group.POST("/register", agent.Register)
	group.PUT("/deregister", agent.DeRegister)
}

func (a *AgentController) GetAgentList(ctx *gin.Context) {
	params := &agentdto.AgentListInput{}
	if err := params.BindValidParm(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	data, err := AgentService.GetAgentList(ctx, params)
	if err != nil {
		log.Logger.Error("查询列表出错", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (a *AgentController) Register(ctx *gin.Context) {
	params := &agentdto.AgentRegisterInput{}
	if err := params.BindValidParm(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	if err := AgentService.Register(ctx, params); err != nil {
		log.Logger.Error("注册失败", err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(ctx, "注册成功")
}

func (a *AgentController) DeRegister(ctx *gin.Context) {
	params := &agentdto.AgentDeRegisterInput{}
	if err := params.BindValidParm(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	if err := AgentService.DeRegister(ctx, params.ServiceName); err != nil {
		log.Logger.Error("注销失败", err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(ctx, "注销成功")
}
