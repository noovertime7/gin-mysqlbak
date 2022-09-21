package dao

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"gorm.io/gorm"
)

type HostDatabase struct {
	Id         int            `json:"host_id" gorm:"primary_key" description:"自增主键"`
	Host       string         `json:"host" gorm:"column:host" description:"任务id"`
	User       string         `json:"user"  gorm:"column:user" description:"是否发送钉钉消息"`
	Password   string         `json:"password"  gorm:"column:password" description:"accessToken"`
	Content    string         `json:"content" gorm:"column:content"`
	HostStatus int            `json:"host_status" gorm:"column:host_status"`
	IsDeleted  int            `json:"is_deleted" gorm:"column:is_deleted"`
	Type       sql.NullInt32  `gorm:"column:type;type:int(11);default:1;comment:1为mysql;2为elastic" json:"type"`
	Avatar     sql.NullString `gorm:"column:avatar;type:varchar(60);comment:头像" json:"avatar"`
}

func (s *HostDatabase) TableName() string {
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

func (s *HostDatabase) UpdatesStatus(tx *gorm.DB) error {
	return tx.Table(s.TableName()).Where("host = ?", s.Host).Updates(map[string]interface{}{
		"host_status": s.HostStatus,
	}).Error
}

func (s *HostDatabase) FindAllHost(tx *gorm.DB) ([]HostDatabase, error) {
	var hostlist []HostDatabase
	if err := tx.Where("is_deleted = 0").Find(&hostlist).Error; err != nil {
		return nil, err
	}
	return hostlist, nil
}

func (s *HostDatabase) PageList(c *gin.Context, tx *gorm.DB, params *dto.HostListInput) ([]HostDatabase, int, error) {
	var total int64 = 0
	var list []HostDatabase
	offset := (params.PageNo - 1) * params.PageSize
	query := tx.WithContext(c)
	query = query.Table(s.TableName()).Where("is_deleted=0")
	query.Find(&list).Count(&total)
	if params.Info != "" {
		query = query.Where("(host like ? )", "%"+params.Info+"%")
	}
	if err := query.Limit(params.PageSize).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return list, int(total), nil
}
