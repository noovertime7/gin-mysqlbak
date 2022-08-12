package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/middleware"
)

type AgentController struct {
}

func AgentRegister(group *gin.RouterGroup) {
	agent := &AgentController{}
	group.GET("/agentlist", agent.GetAgentList)

}

func (a *AgentController) GetAgentList(ctx *gin.Context) {
	out := pkg.GetServiceList()
	middleware.ResponseSuccess(ctx, out)
}
