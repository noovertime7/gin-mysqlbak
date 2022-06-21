package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type BakHistory struct {
	Id        int       `gorm:"primary_key" description:"自增主键"`
	TaskID    int       `gorm:"column:task_id" description:"任务id"`
	Host      string    `gorm:"column:host" description:"主机"`
	DBName    string    `gorm:"column:db_name" description:"库名"`
	BakStatus int       `gorm:"column:bak_status" description:"备份状态"`
	Msg       string    `gorm:"column:message" description:"消息"`
	FileSize  int       `gorm:"column:filesize" description:"文件大小"`
	FileName  string    `gorm:"column:filename" description:"文件名"`
	BakTime   time.Time `gorm:"column:bak_time" description:"备份时间"`
}

func (s *BakHistory) TableName() string {
	return "bak_history"
}

func (s *BakHistory) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(s).Error
}

func (s *BakHistory) Find(c *gin.Context, tx *gorm.DB, search *DingDatabase) (*DingDatabase, error) {
	out := &DingDatabase{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}
