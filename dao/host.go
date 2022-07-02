package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"gorm.io/gorm"
)

type HostDatabase struct {
	Id        int    `gorm:"primary_key" description:"自增主键"`
	Host      string `json:"host" gorm:"column:host" description:"任务id"`
	User      string `json:"user"  gorm:"column:user" description:"是否发送钉钉消息"`
	Password  string `json:"password"  gorm:"column:password" description:"accessToken"`
	IsDeleted int    `json:"is_deleted" gorm:"column:is_deleted"`
}

func (h *HostDatabase) TableName() string {
	return "t_host"
}

func (s *HostDatabase) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(s).Error
}

func (s *HostDatabase) Find(c *gin.Context, tx *gorm.DB, search *HostDatabase) (*HostDatabase, error) {
	out := &HostDatabase{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *HostDatabase) PageList(c *gin.Context, tx *gorm.DB, params *dto.HostListInput) ([]HostDatabase, int, error) {
	var total int64 = 0
	var list []HostDatabase
	offset := (params.PageNo - 1) * params.PageSize
	query := tx.WithContext(c)
	query = query.Table(s.TableName()).Where("is_deleted=0")
	if params.Info != "" {
		query = query.Where("(host like ? )", "%"+params.Info+"%")
	}
	if err := query.Limit(params.PageSize).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	query.Limit(params.PageSize).Offset(offset).Find(&list).Count(&total)
	return list, int(total), nil
}
