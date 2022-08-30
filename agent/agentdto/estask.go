package agentdto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

type ESBakTaskAddInput struct {
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名"  validate:"required"`
	Host        string `json:"host" form:"host" comment:"主机"  validate:"required"`
	User        string `json:"user" form:"user" comment:"用户名"  validate:"required"`
	Password    string `json:"password" form:"password" comment:"密码"  validate:"required"`
	BackupCycle string `json:"backup_cycle" form:"backup_cycle" comment:"备份周期"  validate:"required"`
	KeepNumber  int64  `json:"keep_number" form:"keep_number" comment:"保存时间"  validate:""`
}

type ESBakTaskUpdateInput struct {
	ID          int64  `json:"id" form:"id" comment:"ID"  validate:"required"`
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名"  validate:"required"`
	Host        string `json:"host" form:"host" comment:"主机"  validate:"required"`
	User        string `json:"user" form:"user" comment:"用户名"  validate:"required"`
	Password    string `json:"password" form:"password" comment:"密码"  validate:"required"`
	BackupCycle string `json:"backup_cycle" form:"backup_cycle" comment:"备份周期"  validate:"required"`
	KeepNumber  int64  `json:"keep_number" form:"keep_number" comment:"保存时间"  validate:""`
}

type ESBakTaskIDInput struct {
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名"  validate:"required"`
	ID          int64  `json:"id" form:"id" comment:"ID"  validate:"required"`
}

type ESBakTaskListInput struct {
	ServiceName string `form:"service_name" json:"service_name" comment:"服务名"   validate:"required" example:"test.local"`
	Info        string `form:"info" json:"info" comment:"关键词"   validate:"" example:""`
	PageNo      int64  `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize    int64  `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
}

func (d *ESBakTaskAddInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *ESBakTaskUpdateInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *ESBakTaskIDInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *ESBakTaskListInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}
