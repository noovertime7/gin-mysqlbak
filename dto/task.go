package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

// TaskAddInput 新增task
type TaskAddInput struct {
	HostID          int    `json:"host_id" form:"host_id" validate:"required"`
	DBName          string `form:"db_name" json:"db_name" comment:"库名"   validate:"required" example:"123456"`
	BackupCycle     string `form:"backup_cycle" json:"backup_cycle" comment:"数据库备份时间"   validate:"required,is_valid_bycle" example:"123456"`
	KeepNumber      int    `form:"keep_number" json:"keep_number" comment:"保留周期"   validate:"required" example:"123456"`
	IsAllDBBak      int    `form:"is_all_dbBak" json:"is_all_dbBak" comment:"是否全库备份 0开启 1关闭"  example:"123456"`
	IsDingSend      int    `json:"is_ding_send"`
	DingAccessToken string `json:"ding_access_token"`
	DingSecret      string `json:"ding_secret"`
	OssType         int    `json:"oss_type" validate:"" `
	IsOssSave       int    `json:"is_oss_save" validate:"" `
	Endpoint        string `json:"endpoint" validate:"" `
	OssAccess       string `json:"oss_access" validate:"" `
	OssSecret       string `json:"oss_secret" validate:"" `
	BucketName      string `json:"bucket_name" validate:"" `
	Directory       string `json:"directory" validate:"" `
}

// TaskUpdateInput 更新任务
type TaskUpdateInput struct {
	ID              int    `json:"id" form:"id" validate:"required"`
	HostID          int    `json:"host_id" form:"host_id" validate:"required"`
	DBName          string `form:"db_name" json:"db_name" comment:"库名"   validate:"required" example:"123456"`
	BackupCycle     string `form:"backup_cycle" json:"backup_cycle" comment:"数据库备份时间"   validate:"required,is_valid_bycle" example:"123456"`
	KeepNumber      int    `form:"keep_number" json:"keep_number" comment:"保留周期"   validate:"required" example:"123456"`
	IsAllDBBak      int    `form:"is_all_dbBak" json:"is_all_dbBak" comment:"是否全库备份 0开启 1关闭"  example:"123456"`
	IsDingSend      int    `json:"is_ding_send"`
	DingAccessToken string `json:"ding_access_token"`
	DingSecret      string `json:"ding_secret"`
	OssType         int    `json:"oss_type" validate:"" `
	IsOssSave       int    `json:"is_oss_save" validate:"" `
	Endpoint        string `json:"endpoint" validate:"" `
	OssAccess       string `json:"oss_access" validate:"" `
	OssSecret       string `json:"oss_secret" validate:"" `
	BucketName      string `json:"bucket_name" validate:"" `
	Directory       string `json:"directory" validate:"" `
}

func (d *TaskUpdateInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

// TaskListInput 通过page page-size 查询服务信息
type TaskListInput struct {
	HostId   int    `json:"host_id" form:"host_id" validate:""`
	Info     string `form:"info" json:"info" comment:"关键词"   validate:"" example:""`
	Status   int    `form:"status" json:"status" comment:"任务状态,1为关闭，2为开启" validate:"" example:""`
	PageNo   int    `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize int    `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
}

type TaskListOutput struct {
	Total    int               `form:"total" json:"total" comment:"总数"   validate:"" example:""`
	List     []TaskListOutItem `json:"list" form:"list" comment:"列表" example:"" validate:""` //列表
	PageNo   int               `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize int               `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
}

type TaskListOutItem struct {
	ID          int    `json:"id" form:"id"`
	Host        string `json:"host" form:"host"`
	HostID      int    `json:"host_id" form:"host_id"`
	DBName      string `json:"db_name" form:"db_name"`
	BackupCycle string `json:"backup_cycle" form:"backup_cycle"`
	KeepNumber  int    `json:"keep_number" form:"keep_number"`
	Status      int    `json:"status" form:"status"`
	CreateAt    string `json:"create_at" form:"create_at"`
}

type TaskDeTailOutPut struct {
	ID              int    `json:"id" form:"id"`
	Host            string `json:"host" form:"host"`
	HostID          int    `json:"host_id" form:"host_id"`
	DBName          string `json:"db_name" form:"db_name"`
	BackupCycle     string `json:"backup_cycle" form:"backup_cycle"`
	KeepNumber      int    `json:"keep_number" form:"keep_number"`
	CreateAt        string `json:"create_at" form:"create_at"`
	DingID          int    `json:"ding_id"`
	IsDingSend      int    `json:"is_ding_send"  comment:"是否发送钉钉消息"`
	DingAccessToken string `json:"ding_access_token"  comment:"accessToken"`
	DingSecret      string `json:"ding_secret" comment:"secret"`
	OssID           int    `json:"oss_id"`
	IsOssSave       int    `json:"is_oss_save"  comment:"是否保存到oss中 0关闭1开启"`
	OssType         int    `json:"oss_type"  comment:"oss类型"`
	Endpoint        string `json:"endpoint"   comment:"endpoint"`
	OssAccess       string `json:"oss_access"  comment:"ossaccess"`
	OssSecret       string `json:"oss_secret"   comment:"secret"`
	BucketName      string `json:"bucket_name"  comment:"bucket名字"`
	Directory       string `json:"directory"  comment:"目录"`
}

// TaskDeleteInput 删除task
type TaskDeleteInput struct {
	ID int `json:"id" form:"id" validate:"required"`
}

func (d *TaskDeleteInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *TaskAddInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *TaskListInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}
