package agentdto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

type AgentListInput struct {
	Info      string `form:"info" json:"info" comment:"关键词"   validate:"" example:""`
	PageNo    int    `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize  int    `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
	Status    string `form:"status" json:"status" validate:""`
	SortField string `form:"sortField" json:"sortField" comment:"排序字段" `
	SortOrder string `json:"sortOrder" form:"sortOrder" comment:"排序规则"`
}

type AgentListOutPut struct {
	Total           int                `json:"total"`
	AgentOutPutItem []*AgentOutPutItem `json:"list"`
	PageNo          int                `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize        int                `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
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

type AgentServiceNumInfoOutput struct {
	AllServices    int `json:"all_services"`
	AllTasks       int `json:"all_tasks"`
	AllFinishTasks int `json:"all_finish_tasks"`
}

func (a *AgentRegisterInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}

func (a *AgentListInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}

func (a *AgentDeRegisterInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}
