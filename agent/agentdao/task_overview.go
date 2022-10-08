package agentdao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type TaskOverview struct {
	ID          int64         `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	ServiceName string        `gorm:"column:service_name;type:varchar(255);NOT NULL" json:"service_name"`
	HostId      int64         `gorm:"column:host_id;type:int(11);NOT NULL" json:"host_id"`
	Host        string        `gorm:"column:host;type:varchar(255);NOT NULL" json:"host"`
	TaskId      int64         `gorm:"column:task_id;type:int(11);NOT NULL" json:"task_id"`
	DbName      string        `gorm:"column:db_name;type:varchar(255);NOT NULL" json:"db_name"`
	BackupCycle string        `gorm:"column:backup_cycle;type:varchar(255);NOT NULL" json:"backup_cycle"`
	KeepNumber  int64         `gorm:"column:keep_number;type:int(11);NOT NULL" json:"keep_number"`
	Status      sql.NullInt64 `gorm:"column:status;type:int(11)" json:"status"`
	Type        int64         `gorm:"column:type;type:int(11)" json:"type"`
	CreatedAt   time.Time     `gorm:"column:created_at;type:datetime;NOT NULL" json:"created_at"`
	UpdateAt    time.Time     `gorm:"column:update_at;type:datetime;NOT NULL" json:"update_at"`
	IsDeleted   sql.NullInt64 `gorm:"column:is_deleted;type:int(11);NOT NULL" json:"is_deleted"`
	DeletedAt   time.Time     `json:"deleted_at" gorm:"column:deleted_at" description:"删除时间"`
}

func (t *TaskOverview) TableName() string {
	return "cluster_task_overview"
}

func (t *TaskOverview) Save(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Save(t).Error
}

func (t *TaskOverview) Find(ctx context.Context, tx *gorm.DB, search *TaskOverview) (*TaskOverview, error) {
	out := &TaskOverview{}
	return out, tx.WithContext(ctx).Where(search).Find(out).Error
}

func (t *TaskOverview) Updates(ctx context.Context, tx *gorm.DB, id int64) error {
	return tx.WithContext(ctx).Where("id = ?", id).Updates(t).Error
}

func (t *TaskOverview) Delete(ctx context.Context, tx *gorm.DB) error {
	if t.ID == 0 {
		return errors.New("ID为空")
	}
	return tx.WithContext(ctx).Delete(t).Error
}

func (t *TaskOverview) PageList(c *gin.Context, tx *gorm.DB, params *agentdto.TaskOverViewListInput) ([]TaskOverview, int, error) {
	var total int64 = 0
	var list []TaskOverview
	offset := (params.PageNo - 1) * params.PageSize
	query := tx.WithContext(c)
	query = query.Table(t.TableName())
	switch params.Status {
	//查询关闭状态任务
	case 1:
		query = query.Table(t.TableName()).Where("status != 1")
		if params.Type != 0 {
			query = query.Table(t.TableName()).Where("status != 1 and type = ?", params.Type)
		}
	case 2:
		query = query.Table(t.TableName()).Where("status = 1")
		if params.Type != 0 {
			query = query.Table(t.TableName()).Where("status = 1 and type = ?", params.Type)
		}
	default:
		if params.Type != 0 {
			query = query.Table(t.TableName()).Where(" type = ?", params.Type)
		}
	}
	query.Find(&list).Count(&total)
	if params.Info != "" {
		query = query.Where("( db_name like ? or service_name like ? or host like ? )", "%"+params.Info+"%", "%"+params.Info+"%", "%"+params.Info+"%")
	}
	var sortRules string
	switch params.SortOrder {
	case "descend":
		sortRules = "desc"
	case "ascend":
		sortRules = "asc"
	default:
		sortRules = "desc"
	}
	if params.SortField == "" {
		params.SortField = "id"
		sortRules = "desc"
	}
	if err := query.Limit(int(params.PageSize)).Offset(int(offset)).Order(fmt.Sprintf("%s %s", params.SortField, sortRules)).Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return list, int(total), nil
}
