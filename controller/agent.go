package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/services"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type AgentController struct{}

func AgentRegister(group *gin.RouterGroup) {
	agent := &AgentController{}
	group.POST("/register", agent.Register)
}

func (a *AgentController) Register(c *gin.Context) {
	params := &dto.AgentRegisterInfo{}
	if err := params.BindValidParm(c); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	log.Logger.Infof("agent注册,AgentName:%v,AgentIPPort:%v", params.AgentName, params.AgentIpPort)
	agentservice := services.AgentService{}
	if err := agentservice.AddAgent(c, params); err != nil {
		log.Logger.Error("agent添加失败")
		middleware.ResponseError(c, 20001, err)
		return
	}
	middleware.ResponseSuccess(c, "register success")
}
