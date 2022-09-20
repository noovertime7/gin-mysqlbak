package agentdto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

type StartBakInput struct {
	ServiceName string `json:"service_name" form:"service_name" validate:"required" comment:"服务名"`
	TaskID      int64  `json:"task_id" form:"task_id" validate:"required"`
}
type StopBakInput struct {
	ServiceName string `json:"service_name" form:"service_name" validate:"required" comment:"服务名"`
	TaskID      int64  `json:"task_id" form:"task_id" validate:"required"`
}
type StartBakByHostInput struct {
	ServiceName string `json:"service_name" form:"service_name" validate:"required" comment:"服务名"`
	HostID      int64  `json:"host_id" form:"host_id" validate:"required"`
}
type StopBakByHostInput struct {
	ServiceName string `json:"service_name" form:"service_name" validate:"required" comment:"服务名"`
	HostID      int64  `json:"host_id" form:"host_id" validate:"required"`
}

func (b *StartBakInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, b)
}

func (b *StopBakInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, b)
}

func (b *StartBakByHostInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, b)
}

func (b *StopBakByHostInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, b)
}
