package agentdto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

type ESBakStartInput struct {
	ID          int64  `json:"task_id" form:"task_id" comment:"ID"  validate:"required"`
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名"  validate:"required"`
}

type ESBakStopInput struct {
	ID          int64  `json:"task_id" form:"task_id" comment:"ID"  validate:"required"`
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名"  validate:"required"`
}

func (d *ESBakStopInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *ESBakStartInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}
