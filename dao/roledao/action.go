package roledao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ActionDB struct {
	Id           int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	PermissionId int    `gorm:"column:permission_id;type:int(11)" json:"permission_id"`
	Describe     string `gorm:"column:describe;type:varchar(20)" json:"describe"`
	DefaultCheck int    `gorm:"column:default_check;type:int(11)" json:"default_check"`
	Action       string `gorm:"column:action;type:varchar(20)" json:"action"`
}

func (a *ActionDB) TableName() string {
	return "t_action"
}

func (a *ActionDB) FindActions(ctx *gin.Context, tx *gorm.DB, search *ActionDB) ([]*ActionDB, error) {
	var out []*ActionDB
	if err := tx.WithContext(ctx).Where(search).Find(&out).Error; err != nil {
		return nil, err
	}
	return out, nil
}
