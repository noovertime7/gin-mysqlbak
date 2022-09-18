package dao

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	Id        int           `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName  string        `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Salt      string        `json:"salt" gorm:"column:salt" description:"盐"`
	Password  string        `json:"password" gorm:"column:password" description:"密码"`
	UpdatedAt time.Time     `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time     `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete  sql.NullInt32 `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
	GroupId   int           `gorm:"column:group_id;type:int(11)" json:"group_id"`
	InfoId    int           `gorm:"column:info_id;type:int(11);comment:用户信息关联" json:"info_id"`
	Status    int           `gorm:"column:status;type:int(11);comment:在线状态" json:"status"`
	LoginTime time.Time     `gorm:"column:login_time;type:datetime" json:"login_time"`
}

func (t *Admin) TableName() string {
	return "t_admin"
}

func (t *Admin) LoginCheck(c *gin.Context, tx *gorm.DB, param *dto.AdminLoginInput) (*Admin, error) {
	admininfo, err := t.Find(c, tx, &Admin{UserName: param.UserName, IsDelete: sql.NullInt32{0, true}})
	if err == gorm.ErrRecordNotFound || admininfo.Id == 0 {
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
	return out, tx.WithContext(c).Where(search).Find(out).Error
}

func (t *Admin) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(t).Error
}

func (t *Admin) UpdateStatus(ctx *gin.Context, tx *gorm.DB, info *Admin) error {
	if info.Id == 0 {
		return errors.New("ID 为 0")
	}
	return tx.WithContext(ctx).Table(t.TableName()).Where("id = ?", info.Id).Updates(map[string]interface{}{
		"status":     info.Status,
		"login_time": info.LoginTime,
	}).Error
}

func (a *Admin) Updates(ctx *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Table(a.TableName()).Updates(a).Error
}

func (u *Admin) PageList(c *gin.Context, tx *gorm.DB, params *dto.GroupUserListInput, groupId int) ([]*Admin, int, error) {
	var total int64 = 0
	var list []*Admin
	offset := (params.PageNo - 1) * params.PageSize
	query := tx.WithContext(c)
	if groupId == 0 {
		query = query.Table(u.TableName()).Where("is_delete = 0")
	} else {
		query = query.Table(u.TableName()).Where("is_delete = 0 and group_id = ?", groupId)
	}
	query.Find(&list).Count(&total)
	if params.Info != "" {
		query = query.Where("(user_name like ?)", "%"+params.Info+"%")
	}
	if err := query.Limit(params.PageSize).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return list, int(total), nil
}
