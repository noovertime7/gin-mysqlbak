package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

// HostAddInput 添加
type HostAddInput struct {
	Host     string `form:"host" json:"host" comment:"数据库备份主机地址加端口"   validate:"required,host_valid" example:"127.0.0.1"`
	User     string `form:"username" json:"username" comment:"用户"   validate:"required" example:"123456"`
	Content  string `form:"content" json:"content" comment:"备注"   validate:"" example:"123456"`
	Password string `form:"password" json:"password" comment:"数据库密码"   validate:"required" example:"123456"`
}

func (d *HostAddInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

// HostDeleteInput  删除
type HostDeleteInput struct {
	ID int `json:"id" form:"id" validate:"required"`
}

func (d *HostDeleteInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

// HostUpdateInput 修改
type HostUpdateInput struct {
	ID       int    `json:"id" form:"id" validate:"required"`
	Host     string `form:"host" json:"host" comment:"数据库备份主机地址加端口"   validate:"required,host_valid" example:"127.0.0.1"`
	User     string `form:"username" json:"username" comment:"用户"   validate:"required" example:"123456"`
	Content  string `form:"content" json:"content" comment:"备注"   validate:"" example:"123456"`
	Password string `form:"password" json:"password" comment:"数据库密码"   validate:"required" example:"123456"`
}

func (d *HostUpdateInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

// HostListInput 查询
type HostListInput struct {
	Info     string `form:"info" json:"info" comment:"关键词"   validate:"" example:""`
	PageNo   int    `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize int    `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
}

func (d *HostListInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

type HostListOutput struct {
	Total int               `form:"total" json:"total" comment:"总数"   validate:"" example:""`
	List  []HostListOutItem `json:"list" form:"list" comment:"列表" example:"" validate:""` //列表
}

type HostListOutItem struct {
	ID         int    `json:"id" form:"id"`
	Host       string `json:"host" form:"host"`
	User       string `json:"username" comment:"用户"`
	Password   string `json:"password" comment:"数据库密码"`
	HostStatus int    `json:"host_status"`
	Content    string `json:"content"`
	TaskNum    int    `json:"task_num"`
}

type HostIDInput struct {
	HostID int `json:"host_id" form:"host_id"`
}

func (d *HostIDInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}
