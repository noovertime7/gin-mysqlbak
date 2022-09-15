package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	Id        int       `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName  string    `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Salt      string    `json:"salt" gorm:"column:salt" description:"盐"`
	Password  string    `json:"password" gorm:"column:password" description:"密码"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete  int       `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
	Role      int       `gorm:"column:role;type:int(11)" json:"role"`
	Status    int       `gorm:"column:status;type:int(11);comment:在线状态" json:"status"`
}

func (t *Admin) TableName() string {
	return "t_admin"
}

func (t *Admin) LoginCheck(c *gin.Context, tx *gorm.DB, param *dto.AdminLoginInput) (*Admin, error) {
	admininfo, err := t.Find(c, tx, &Admin{UserName: param.UserName, IsDelete: 0})
	if err != nil {
		return nil, errors.New("用户信息不存在")
	}
	saltPassword := public.GenSaltPassword(admininfo.Salt, param.Password)
	if admininfo.Password != saltPassword {
		return nil, errors.New("密码错误请重新输入")
	}
	return admininfo, nil
}

func (t *Admin) Find(c *gin.Context, tx *gorm.DB, search *Admin) (*Admin, error) {
	out := &Admin{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *Admin) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(t).Error
}

func (t *Admin) UpdateStatus(ctx *gin.Context, tx *gorm.DB, info *Admin) error {
	if info.Id == 0 {
		return errors.New("ID 为 0")
	}
	return tx.WithContext(ctx).Table(t.TableName()).Where("id = ?", info.Id).Updates(map[string]interface{}{
		"status": info.Status,
	}).Error
}
