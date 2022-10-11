package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

type PanelGroupDataOutPut struct {
	TaskNum       int `json:"task_num"`
	HistoryNum    int `json:"history_num"`
	RunningProNum int `json:"running_pro_num"`
	HostNum       int `json:"host_num"`
}

type DashServiceStatOutput struct {
	Legend []string                    `json:"legend"`
	Data   []DashServiceStatItemOutput `json:"data"`
}

type DashServiceStatItemOutput struct {
	HostID int    `json:"host_id"`
	Name   string `json:"name"`
	Value  int64  `json:"value"`
}

type ServiceTaskOutPut struct {
	ServiceName string `json:"item"`
	TaskNum     int64  `json:"count"`
}

type AgentDateInfoInput struct {
	Day int `json:"day" form:"day"`
}

type DateInfoOut struct {
	TaskTotal         int64              `json:"task_total"`
	FinishTotal       int64              `json:"finish_total"`
	TaskIncreaseNum   int64              `json:"task_increase_num"`
	TaskDecreaseNum   int64              `json:"task_decrease_num"`
	FinishIncreaseNum int64              `json:"finish_increase_num"`
	FinishDecreaseNum int64              `json:"finish_decrease_num"`
	List              []*DateInfoOutItem `json:"list"`
}

type DateInfoOutItem struct {
	Date      string `json:"date"`
	AgentNum  int64  `json:"agent_num"`
	TaskNum   int64  `json:"task_num"`
	FinishNum int64  `json:"finish_num"`
}

func (d *AgentDateInfoInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}
