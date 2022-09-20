package dao

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/public"
	"gorm.io/gorm"
	"time"
)

type BakHistory struct {
	Id         int           `gorm:"primary_key" description:"自增主键"`
	TaskID     int           `gorm:"column:task_id" description:"任务id"`
	Host       string        `gorm:"column:host" description:"主机"`
	DBName     string        `gorm:"column:db_name" description:"库名"`
	OssStatus  int           `gorm:"column:oss_status"  description:"钉钉发送状态"`
	DingStatus int           `gorm:"column:ding_status"  description:"OSS保存状态"`
	BakStatus  int           `gorm:"column:bak_status" description:"备份状态"`
	Msg        string        `gorm:"column:message" description:"消息"`
	FileSize   int           `gorm:"column:file_size" description:"文件大小"`
	FileName   string        `gorm:"column:file_name" description:"文件名"`
	BakTime    time.Time     `gorm:"column:bak_time" description:"备份时间"`
	IsDelete   sql.NullInt32 `gorm:"column:is_delete;type:int(11);NOT NULL" json:"is_delete"`
}

func (b *BakHistory) TableName() string {
	return "bak_history"
}

func (b *BakHistory) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(b).Error
}

func (b *BakHistory) Updates(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Table(b.TableName()).Updates(b).Error
}

func (b *BakHistory) Find(c *gin.Context, tx *gorm.DB, search *BakHistory) (*BakHistory, error) {
	out := &BakHistory{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FindAllHistory 查询所有备份历史记录
func (b *BakHistory) FindAllHistory(c *gin.Context, tx *gorm.DB, status string) ([]*BakHistory, error) {
	var result []*BakHistory
	switch status {
	case "":
		err := tx.WithContext(c).Order("id desc").Where("is_delete = 0").Find(&result).Error
		if err != nil {
			return nil, err
		}
		return result, nil
	case "success":
		err := tx.WithContext(c).Where("is_delete = 0 and bak_status = ?", 1).Order("id desc").Find(&result).Error
		if err != nil {
			return nil, err
		}
		return result, nil
	case "fail":
		err := tx.WithContext(c).Where("is_delete = 0 and bak_status = ?", 0).Order("id desc").Find(&result).Error
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, nil
}

// FindByDate 查询7天内数据
func (b *BakHistory) FindByDate(ctx *gin.Context, tx *gorm.DB, num int) ([]BakHistory, error) {
	var out []BakHistory
	return out, tx.WithContext(ctx).Raw("SELECT * FROM bak_history WHERE date_sub(curdate(), interval ? day) <= date(bak_time);", num).Scan(&out).Error
}

func (b *BakHistory) PageList(c *gin.Context, tx *gorm.DB, params *dto.HistoryListInput) ([]BakHistory, int, error) {
	var total int64 = 0
	var list []BakHistory
	offset := (params.PageNo - 1) * params.PageSize
	query := tx.WithContext(c)
	query = query.Table(b.TableName()).Where("is_delete = 0")
	query.Find(&list).Count(&total)
	if params.Status != "" {
		switch params.Status {
		case public.HistoryStatusAll:
			if params.Info != "" {
				searchInfo := "%" + params.Info
				query = query.Where(fmt.Sprintf(" host like '%%%s' or db_name like'%%%s'  ", searchInfo, searchInfo))
			} else {
				query = query.Table(b.TableName()).Where("is_delete = 0")
			}
		case public.HistoryStatusSuccess:
			if params.Info != "" {
				searchInfo := "%" + params.Info
				query = query.Where("(host like ? or db_name like ?)", searchInfo, searchInfo)
			} else {
				query = query.Where("message = 'success' ")
			}
		case public.HistoryStatusFail:
			if params.Info != "" {
				searchInfo := "%" + params.Info
				query = query.Where("host like ? or db_name like ?", searchInfo, searchInfo)
			} else {
				query = query.Where("message != 'success' ")
			}
		}
	}
	var sortRules string
	switch params.SortOrder {
	case "descend":
		sortRules = "desc"
	case "ascend":
		sortRules = "asc"
	}
	if params.SortField == "" {
		params.SortField = "id"
	}
	if err := query.Limit(params.PageSize).Offset(offset).Order(fmt.Sprintf("%s %s", params.SortField, sortRules)).Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return list, int(total), nil
}
