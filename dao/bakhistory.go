package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"gorm.io/gorm"
	"time"
)

type BakHistory struct {
	Id         int       `gorm:"primary_key" description:"自增主键"`
	TaskID     int       `gorm:"column:task_id" description:"任务id"`
	Host       string    `gorm:"column:host" description:"主机"`
	DBName     string    `gorm:"column:db_name" description:"库名"`
	OssStatus  int       `gorm:"column:oss_status"  description:"钉钉发送状态"`
	DingStatus int       `gorm:"column:ding_status"  description:"OSS保存状态"`
	BakStatus  int       `gorm:"column:bak_status" description:"备份状态"`
	Msg        string    `gorm:"column:message" description:"消息"`
	FileSize   int       `gorm:"column:filesize" description:"文件大小"`
	FileName   string    `gorm:"column:filename" description:"文件名"`
	BakTime    time.Time `gorm:"column:bak_time" description:"备份时间"`
}

func (s *BakHistory) TableName() string {
	return "bak_history"
}

func (s *BakHistory) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(s).Error
}

func (s *BakHistory) Find(c *gin.Context, tx *gorm.DB, search *BakHistory) (*BakHistory, error) {
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
		err := tx.WithContext(c).Order("id desc").Find(&result).Error
		if err != nil {
			return nil, err
		}
		return result, nil
	case "success":
		err := tx.WithContext(c).Where("bak_status = ?", 1).Order("id desc").Find(&result).Error
		if err != nil {
			return nil, err
		}
		return result, nil
	case "fail":
		err := tx.WithContext(c).Where("bak_status = ?", 0).Order("id desc").Find(&result).Error
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, nil
}

func (s *BakHistory) PageList(c *gin.Context, tx *gorm.DB, params *dto.HistoryListInput) ([]BakHistory, int, error) {
	var total int64 = 0
	list := []BakHistory{}
	offset := (params.PageNo - 1) * params.PageSize
	query := tx.WithContext(c)
	query = query.Table(s.TableName())
	query.Find(&list).Count(&total)
	if params.Info != "" {
		query = query.Where("(host like ? or db_name like ?)", "%"+params.Info+"%", "%"+params.Info+"%")
	}
	if params.Sort == "aesc" {
		if err := query.Limit(params.PageSize).Offset(offset).Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0, err
		}
	} else {
		if err := query.Limit(params.PageSize).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0, err
		}
	}
	return list, int(total), nil
}
