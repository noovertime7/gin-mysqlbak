package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

type Bak struct {
	ID int `json:"id" form:"id" validate:"required"`
}

func (d *Bak) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

type BakHistoryOutPut struct {
	Host    string `json:"host"`
	DBName  string `json:"db_name"`
	Message string `json:"message"`
	Baktime string `json:"bak_time"`
}

type HistoryIDInput struct {
	ID int `json:"id" form:"id"`
}

type HistoryListInput struct {
	Info      string `form:"info" json:"info" comment:"关键词"   validate:"" example:""`
	PageNo    int    `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize  int    `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
	Status    string `form:"status" json:"status" validate:""`
	SortField string `form:"sortField" json:"sortField" comment:"排序字段" `
	SortOrder string `json:"sortOrder" form:"sortOrder" comment:"排序规则"`
}

type HistoryListOutput struct {
	Total    int                  `form:"total" json:"total" comment:"总数"   validate:"" example:""`
	List     []HistoryListOutItem `json:"list" form:"list" comment:"列表" example:"" validate:""` //列表
	PageNo   int                  `form:"page_no" json:"page_no" comment:"当前页数"   validate:"" example:"1"`
	PageSize int                  `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
}

type HistoryListOutItem struct {
	ID         int    `json:"id"`
	Host       string `json:"host" form:"host"`
	DBName     string `json:"db_name" form:"db_name"`
	DingStatus int    `json:"ding_status" form:"ding_status"`
	OSSStatus  int    `json:"oss_status" form:"oss_status"`
	Message    string `json:"message" form:"message"`
	FileSize   int    `json:"file_size" form:"file_size"`
	FileName   string `json:"file_name" form:"file_name"`
	BakTime    string `json:"bak_time" form:"bak_time"`
}

type HistoryNumInfoOutput struct {
	WeekNums    int    `json:"week_nums"`
	AllNums     int    `json:"all_nums"`
	AllFileSize string `json:"all_filesize"`
}

func (d *HistoryListInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}

func (d *HistoryIDInput) BindValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}
