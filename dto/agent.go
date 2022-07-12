package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

type AgentRegisterInfo struct {
	AgentName   string `json:"agent_name"`
	AgentIpPort string `json:"agent_ip_port"`
}

func (d *AgentRegisterInfo) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}
