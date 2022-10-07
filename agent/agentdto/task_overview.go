package agentdto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

type TaskOverViewListInput struct {
	Info     string `form:"info" json:"info" comment:"关键词"   validate:"" example:""`
	PageNo   int64  `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize int64  `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
	Type     int64  `form:"type" json:"type"`
	Status   int64  `form:"status" json:"status"`
}

type TaskOverViewListOut struct {
	Total    int64                      `form:"total" json:"total" comment:"总数"   validate:"" example:""`
	List     []*TaskOverViewListOutItem `json:"list" form:"list" comment:"列表" example:"" validate:""` //列表
	PageNo   int64                      `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize int64                      `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
}

type TaskOverViewListOutItem struct {
	ID          int64  `json:"id" form:"db_name"`
	ServiceName string `json:"service_name" form:"db_name"`
	HostID      int64  `json:"host_id" form:"db_name"`
	Host        string `json:"host" form:"db_name"`
	TaskID      int64  `json:"task_id" form:"db_name"`
	DBName      string `json:"db_name" form:"db_name"`
	BackupCycle string `json:"backup_cycle" form:"backup_cycle"`
	KeepNumber  int64  `json:"keep_number" form:"db_name"`
	Status      int64  `json:"status" form:"db_name"`
	Type        int64  `json:"type"`
	IsDeleted   int64  `json:"is_deleted"`
}

type StartOverViewBakInput struct {
	ServiceName string `json:"service_name" form:"service_name" validate:"required" comment:"服务名"`
	TaskID      int64  `json:"task_id" form:"task_id" validate:"required"`
}
type StopOverViewBakInput struct {
	ServiceName string `json:"service_name" form:"service_name" validate:"required" comment:"服务名"`
	TaskID      int64  `json:"task_id" form:"task_id" validate:"required"`
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
