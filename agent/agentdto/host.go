package agentdto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

// HostAddInput 添加
type HostAddInput struct {
	ServiceName string `form:"service_name" json:"service_name" comment:"服务名"   validate:"required" example:"123456"`
	Host        string `form:"host" json:"host" comment:"数据库备份主机地址加端口"   validate:"required,host_valid" example:"127.0.0.1"`
	User        string `form:"username" json:"username" comment:"用户"   validate:"required" example:"123456"`
	Content     string `form:"content" json:"content" comment:"备注"   validate:"" example:"123456"`
	Password    string `form:"password" json:"password" comment:"数据库密码"   validate:"required" example:"123456"`
	Type        int64  `form:"type" json:"type"  validate:"required"`
}

// HostDeleteInput  删除
type HostDeleteInput struct {
	ServiceName string `form:"service_name" json:"service_name" comment:"服务名"   validate:"required" example:"123456"`
	ID          int    `json:"id" form:"id" validate:"required"`
}

// HostUpdateInput 修改
type HostUpdateInput struct {
	ServiceName string `form:"service_name" json:"service_name" comment:"服务名"   validate:"required" example:"123456"`
	ID          int    `json:"id" form:"id" validate:"required"`
	Host        string `form:"host" json:"host" comment:"数据库备份主机地址加端口"   validate:"required,host_valid" example:"127.0.0.1"`
	User        string `form:"username" json:"username" comment:"用户"   validate:"required" example:"123456"`
	Content     string `form:"content" json:"content" comment:"备注"   validate:"" example:"123456"`
	Password    string `form:"password" json:"password" comment:"数据库密码"   validate:"required" example:"123456"`
	Type        int64  `form:"type" json:"type"  validate:"required"`
}

// HostListInput 查询
type HostListInput struct {
	ServiceName string `form:"service_name" json:"service_name" comment:"服务名"   validate:"required" example:"test5.local"`
	Info        string `form:"info" json:"info" comment:"关键词"   validate:"" example:""`
	PageNo      int64  `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize    int64  `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
}

type HostIDInput struct {
	ServiceName string `form:"service_name" json:"service_name" comment:"服务名"   validate:"required" example:"test5.local"`
	HostID      int64  `json:"host_id" form:"host_id" validate:"required"`
}

type HostNamesInput struct {
	Type        int64  `form:"type" json:"type" comment:"服务类型"  validate:"required"`
	ServiceName string `form:"service_name" json:"service_name" comment:"服务名"   validate:"required" example:"test5.local"`
}

func (d *HostNamesInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *HostUpdateInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *HostListInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *HostAddInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *HostDeleteInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *HostIDInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}
