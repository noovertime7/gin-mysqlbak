package agentdto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

type AgentListOutPut struct {
	Total           int                `json:"total"`
	AgentOutPutItem []*AgentOutPutItem `json:"agent_list"`
}

type AgentOutPutItem struct {
	ServiceName string `json:"service_name"`
	Address     string `json:"address"`
	Content     string `json:"content"`
	AgentStatus int    `json:"agent_status"`
	LastTime    string `json:"last_time"`
	TaskNum     int    `json:"task_num"`
	FinishNum   int    `json:"finish_num"`
	CreateAt    string `json:"create_at"`
}

type AgentRegisterInput struct {
	ServiceName string `json:"service_name"`
	Address     string `json:"address"`
	Content     string `json:"content"`
	TaskNum     int    `json:"task_num"`
	FinishNum   int    `json:"finish_num"`
}
type AgentDeRegisterInput struct {
	ServiceName string `json:"service_name" form:"service_name"`
}

type AgentListInput struct {
	Info     string `form:"info" json:"info" comment:"关键词"   validate:"" example:""`
	PageNo   int    `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize int    `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
}

func (a *AgentRegisterInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}

func (a *AgentListInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}

func (a *AgentDeRegisterInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}
