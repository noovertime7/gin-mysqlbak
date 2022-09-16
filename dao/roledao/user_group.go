package roledao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserGroupDB 用户组表
type UserGroupDB struct {
	Id        int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	RoleId    int    `gorm:"column:role_id;type:int(11)" json:"role_id"`
	GroupName string `gorm:"column:group_name;type:varchar(20)" json:"group_name"`
	Key       string `gorm:"column:key;type:varchar(20)" json:"key"`
}

func (u *UserGroupDB) TableName() string {
	return "t_user_group"
}

func (u *UserGroupDB) Find(ctx *gin.Context, tx *gorm.DB, search *UserGroupDB) (*UserGroupDB, error) {
	out := &UserGroupDB{}
	return out, tx.WithContext(ctx).Where(search).Find(out).Error
}
