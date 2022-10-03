package agentdto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

type EsHistoryIDInput struct {
	ServiceName string `form:"service_name" json:"service_name" comment:"服务名"   validate:"required" example:"test.local"`
	ID          int64  `json:"id" form:"id" validate:"required"`
}

type EsHistorySvcName struct {
	ServiceName string `form:"service_name" json:"service_name" comment:"服务名"   validate:"required" example:"test.local"`
}

type ESHistoryListInput struct {
	ServiceName string `form:"service_name" json:"service_name" comment:"服务名"   validate:"required" example:"test.local"`
	Info        string `form:"info" json:"info" comment:"关键词"   validate:"" example:""`
	PageNo      int64  `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize    int64  `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
	Status      string `form:"status" json:"status" validate:""`
	SortField   string `form:"sortField" json:"sortField" comment:"排序字段" `
	SortOrder   string `json:"sortOrder" form:"sortOrder" comment:"排序规则"`
}

func (d *ESHistoryListInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *EsHistoryIDInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *EsHistorySvcName) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}
