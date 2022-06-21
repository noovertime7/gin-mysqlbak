package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DingDatabase struct {
	Id              int    `gorm:"primary_key" description:"自增主键"`
	TaskID          int    `gorm:"column:task_id" description:"任务id"`
	IsDingSend      int    `gorm:"column:is_ding_send" description:"是否发送钉钉消息"`
	DingAccessToken string `gorm:"column:ding_access_token" description:"accessToken"`
	DingSecret      string `gorm:"column:ding_secret" description:"secret"`
}

func (s *DingDatabase) TableName() string {
	return "t_ding"
}

func (s *DingDatabase) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(s).Error
}

func (d *DingDatabase) Updates(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Where("task_id = ?", d.TaskID).Updates(d).Error
}

func (s *DingDatabase) Find(c *gin.Context, tx *gorm.DB, search *DingDatabase) (*DingDatabase, error) {
	out := &DingDatabase{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}
