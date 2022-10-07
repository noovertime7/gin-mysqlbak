package agentdao

import (
	"context"
	"database/sql"
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
