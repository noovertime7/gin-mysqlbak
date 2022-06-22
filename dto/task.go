package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
	"time"
)

// TaskAddInput 新增task
type TaskAddInput struct {
	Host            string `form:"host" json:"host" comment:"数据库备份主机地址加端口"   validate:"required" example:"127.0.0.1"`
	Password        string `form:"password" json:"password" comment:"数据库密码"   validate:"required" example:"123456"`
	User            string `form:"user" json:"user" comment:"用户"   validate:"required" example:"123456"`
	DBName          string `form:"db_name" json:"db_name" comment:"库名"   validate:"required" example:"123456"`
	BackupCycle     string `form:"backup_cycle" json:"backup_cycle" comment:"备份时间"   validate:"required" example:"123456"`
	KeepNumber      int    `form:"keep_number" json:"keep_number" comment:"保留周期"   validate:"required" example:"123456"`
	IsAllDBBak      int    `form:"is_all_dbBak" json:"is_all_dbBak" comment:"是否全库备份 0开启 1关闭"  example:"123456"`
	IsDingSend      int    `json:"is_ding_send"`
	DingAccessToken string `json:"ding_access_token"`
	DingSecret      string `json:"ding_secret"`
	OssType         int    `json:"oss_type" validate:"required" `
	IsOssSave       int    `json:"is_oss_save" validate:"required" `
	Endpoint        string `json:"endpoint" validate:"required" `
	OssAccess       string `json:"oss_access" validate:"required" `
	OssSecret       string `json:"oss_secret" validate:"required" `
	BucketName      string `json:"bucket_name" validate:"required" `
	Directory       string `json:"directory" validate:"required" `
}

// TaskUpdateInput 更新任务
type TaskUpdateInput struct {
	ID              int    `json:"id" form:"id" validate:"required"`
	Host            string `form:"host" json:"host" comment:"数据库备份主机地址加端口"   validate:"required" example:"127.0.0.1"`
	Password        string `form:"password" json:"password" comment:"数据库密码"   validate:"required" example:"123456"`
	User            string `form:"user" json:"user" comment:"用户"   validate:"required" example:"123456"`
	DBName          string `form:"db_name" json:"db_name" comment:"库名"   validate:"required" example:"123456"`
	BackupCycle     string `form:"backup_cycle" json:"backup_cycle" comment:"备份时间"   validate:"required" example:"123456"`
	KeepNumber      int    `form:"keep_number" json:"keep_number" comment:"保留周期"   validate:"required" example:"123456"`
	IsAllDBBak      int    `form:"is_all_dbBak" json:"is_all_dbBak" comment:"是否全库备份 0开启 1关闭"  example:"123456"`
	IsDingSend      int    `json:"is_ding_send"`
	DingAccessToken string `json:"ding_access_token"`
	DingSecret      string `json:"ding_secret"`
	OssType         int    `json:"oss_type" validate:"required" `
	IsOssSave       int    `json:"is_oss_save" validate:"required" `
	Endpoint        string `json:"endpoint" validate:"required" `
	OssAccess       string `json:"oss_access" validate:"required" `
	OssSecret       string `json:"oss_secret" validate:"required" `
	BucketName      string `json:"bucket_name" validate:"required" `
	Directory       string `json:"directory" validate:"required" `
}

func (d *TaskUpdateInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

// TaskListInput 通过page pagesize 查询服务信息
type TaskListInput struct {
	Info     string `form:"info" json:"info" comment:"关键词"   validate:"" example:""`
	PageNo   int    `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize int    `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
}

type TaskListOutput struct {
	Total int               `form:"total" json:"total" comment:"总数"   validate:"" example:""`
	List  []TaskListOutItem `json:"list" form:"list" comment:"列表" example:"" validate:""` //列表
}

type TaskListOutItem struct {
	ID          int    `json:"id" form:"id"`
	Host        string `json:"host" form:"id"`
	DBName      string `json:"db_name" form:"id"`
	BackupCycle string `json:"backup_cycle" form:"id"`
	KeepNumber  int    `json:"keep_number" form:"id"`
	Status      bool
	CreateAt    time.Time `json:"create_at" form:"id"`
}

//删除task
type TaskDeleteInput struct {
	ID int `json:"id" form:"id" validate:"required"`
}

func (d *TaskDeleteInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *TaskAddInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *TaskListInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}
