package agentdto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

type BakHistoryOutPut struct {
	Host    string `json:"host"`
	DBName  string `json:"db_name"`
	Message string `json:"message"`
	Baktime string `json:"bak_time"`
}

type BakHistoryDeleteInput struct {
	ServiceName string `form:"service_name" json:"service_name" comment:"服务名"   validate:"required" example:"test.local"`
	ID          int64  `json:"id" form:"id" validate:"required"`
}

type HistoryServiceNameInput struct {
	ServiceName string `form:"service_name" json:"service_name" comment:"服务名"   validate:"required" example:"test.local"`
}

type HistoryListInput struct {
	ServiceName string `form:"service_name" json:"service_name" comment:"服务名"   validate:"required" example:"test.local"`
	Info        string `form:"info" json:"info" comment:"关键词"   validate:"" example:""`
	PageNo      int64  `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize    int64  `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
	Status      string `form:"status" json:"status" validate:""`
	SortField   string `form:"sortField" json:"sortField" comment:"排序字段" `
	SortOrder   string `json:"sortOrder" form:"sortOrder" comment:"排序规则"`
}

type HistoryListOutput struct {
	Total int64                `form:"total" json:"total" comment:"总数"   validate:"" example:""`
	List  []HistoryListOutItem `json:"list" form:"list" comment:"列表" example:"" validate:""` //列表
}

type HistoryListOutItem struct {
	ID         int64  `json:"id"`
	Host       string `json:"host" form:"host"`
	DBName     string `json:"db_name" form:"db_name"`
	DingStatus int64  `json:"ding_status" form:"ding_status"`
	OSSStatus  int64  `json:"oss_status" form:"oss_status"`
	Message    string `json:"message" form:"message"`
	FileSize   int64  `json:"file_size" form:"file_size"`
	FileName   string `json:"file_name" form:"file_name"`
	BakTime    string `json:"bak_time" form:"bak_time"`
}

func (d *HistoryListInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *HistoryServiceNameInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *BakHistoryDeleteInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}
