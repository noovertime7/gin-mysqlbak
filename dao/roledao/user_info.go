package roledao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserInfo 用户信息表
type UserInfo struct {
	Id           int    `gorm:"column:id;type:int(11);AUTO_INCREMENT;primary_key" json:"id"`
	Uid          int    `gorm:"column:uid;type:int(11)" json:"uid"`
	Avatar       string `gorm:"column:avatar;type:varchar(255);comment:头像" json:"avatar"`
	CreateId     string `gorm:"column:create_id;type:varchar(20);comment:创建用户" json:"create_id"`
	Email        string `gorm:"column:email;type:varchar(30);comment:邮箱" json:"email"`
	Mobile       string `gorm:"column:mobile;type:varchar(11);comment:手机号" json:"mobile"`
	Nickname     string `gorm:"column:nickname;type:varchar(20);comment:别名" json:"nickname"`
	Introduction string `gorm:"column:introduction;type:varchar(255);comment:介绍" json:"introduction"`
}

func (u *UserInfo) TableName() string {
	return "user_info"
}

func (u *UserInfo) Save(ctx *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Table(u.TableName()).Save(u).Error
}

func (u *UserInfo) Updates(ctx *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Table(u.TableName()).Updates(u).Error
}

func (u *UserInfo) Find(ctx *gin.Context, tx *gorm.DB, search *UserInfo) (*UserInfo, error) {
	out := &UserInfo{}
	return out, tx.WithContext(ctx).Where(search).Find(out).Error
}
