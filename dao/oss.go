package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OssDatabase struct {
	Id         int    `gorm:"primary_key" description:"自增主键"`
	TaskID     int    `gorm:"column:task_id" description:"任务id"`
	IsOssSave  int    `gorm:"column:is_oss_save" description:"是否保存到oss中 0关闭1开启"`
	OssType    int    `gorm:"column:oss_type" description:"oss类型"`
	Endpoint   string `gorm:"column:endpoint" description:"endpoint"`
	OssAccess  string `gorm:"column:oss_access" description:"ossaccess"`
	OssSecret  string `gorm:"column:oss_secret" description:"secret"`
	BucketName string `gorm:"column:bucket_name" description:"bucket名字"`
	Directory  string `gorm:"column:directory" description:"目录"`
}

func (s *OssDatabase) TableName() string {
	return "t_oss"
}

func (s *OssDatabase) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(s).Error
}

func (s *OssDatabase) Find(c *gin.Context, tx *gorm.DB, search *OssDatabase) (*OssDatabase, error) {
	out := &OssDatabase{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *OssDatabase) Updates(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Where("task_id = ?", d.TaskID).Updates(d).Error
}
