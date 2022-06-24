package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OssDatabase struct {
	Id         int    `gorm:"primary_key" description:"自增主键"`
	TaskID     int    `json:"task_id" gorm:"column:task_id" description:"任务id"`
	IsOssSave  int    `json:"is_oss_save" gorm:"column:is_oss_save" description:"是否保存到oss中 0关闭1开启"`
	OssType    int    `json:"oss_type" gorm:"column:oss_type" description:"oss类型"`
	Endpoint   string `json:"endpoint"  gorm:"column:endpoint" description:"endpoint"`
	OssAccess  string `json:"oss_access"  gorm:"column:oss_access" description:"ossaccess"`
	OssSecret  string `json:"oss_secret"  gorm:"column:oss_secret" description:"secret"`
	BucketName string `json:"bucket_name"  gorm:"column:bucket_name" description:"bucket名字"`
	Directory  string `json:"directory" gorm:"column:directory" description:"目录"`
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
	return tx.WithContext(c).Table(d.TableName()).Where("task_id = ?", d.TaskID).Updates(map[string]interface{}{
		"is_oss_save": d.IsOssSave,
		"oss_type":    d.OssType,
		"endpoint":    d.Endpoint,
		"oss_access":  d.OssAccess,
		"oss_secret":  d.OssSecret,
		"bucket_name": d.BucketName,
		"directory":   d.Directory,
	}).Error
}
