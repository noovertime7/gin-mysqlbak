package agentdto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

type TaskOverViewListInput struct {
	Info      string `form:"info" json:"info" comment:"关键词"   validate:"" example:""`
	PageNo    int64  `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize  int64  `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
	Type      int64  `form:"type" json:"type"`
	Status    int64  `form:"status" json:"status"`
	SortField string `form:"sortField" json:"sortField" comment:"排序字段" `
	SortOrder string `json:"sortOrder" form:"sortOrder" comment:"排序规则"`
}

type TaskOverViewListOut struct {
	Total    int64                      `form:"total" json:"total" comment:"总数"   validate:"" example:""`
	List     []*TaskOverViewListOutItem `json:"list" form:"list" comment:"列表" example:"" validate:""` //列表
	PageNo   int64                      `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize int64                      `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
}

type TaskOverViewListOutItem struct {
	ID          int64  `json:"id" form:"id"`
	ServiceName string `json:"service_name" form:"service_name"`
	HostID      int64  `json:"host_id" form:"host_id"`
	Host        string `json:"host" form:"host"`
	TaskID      int64  `json:"task_id" form:"task_id"`
	DBName      string `json:"db_name" form:"db_name"`
	BackupCycle string `json:"backup_cycle" form:"backup_cycle"`
	KeepNumber  int64  `json:"keep_number" form:"keep_number"`
	Status      int64  `json:"status" form:"status"`
	FinishNum   int64  `json:"finish_num"`
	Type        int64  `json:"type"`
	IsDeleted   int64  `json:"is_deleted"`
}

type StartOverViewBakInput struct {
	ID          int64  `json:"id" form:"id" validate:"required"`
	ServiceName string `json:"service_name" form:"service_name" validate:"required" comment:"服务名"`
	TaskID      int64  `json:"task_id" form:"task_id" validate:"required"`
	Type        int64  `json:"type" form:"type" validate:"required"`
}
type StopOverViewBakInput struct {
	ID          int64  `json:"id" form:"id" validate:"required"`
	ServiceName string `json:"service_name" form:"service_name" validate:"required" comment:"服务名"`
	TaskID      int64  `json:"task_id" form:"task_id" validate:"required"`
	Type        int64  `json:"type" form:"type" validate:"required"`
}

type DeleteOverViewTaskInput struct {
	ID          int64  `json:"id" form:"id" validate:"required"`
	ServiceName string `json:"service_name" form:"service_name" validate:"required" comment:"服务名"`
	TaskID      int64  `json:"task_id" form:"task_id" validate:"required"`
	Type        int64  `json:"type" form:"type" validate:"required"`
}

func (d *DeleteOverViewTaskInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *TaskOverViewListInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *StartOverViewBakInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *StopOverViewBakInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}
