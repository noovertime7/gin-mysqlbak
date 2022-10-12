package dao

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type JobHistory struct {
	ID         int           `gorm:"column:id;type:int(11);AUTO_INCREMENT;primary_key" json:"id"`
	JobType    int           `gorm:"column:job_type;type:int(11);NOT NULL" json:"job_type"`
	JobCycle   string        `gorm:"column:job_cycle;type:varchar(20);NOT NULL" json:"job_cycle"`
	Affected   int           `gorm:"column:affected;type:int(11)" json:"affected"`
	Status     sql.NullInt64 `gorm:"column:status;type:int(11);NOT NULL" json:"status"`
	Message    string        `gorm:"column:message;type:varchar(255)" json:"message"`
	UpdateTime time.Time     `gorm:"column:update_time;type:datetime" json:"update_time"`
}

func (j *JobHistory) TableName() string {
	return "job_history"
}

func (j *JobHistory) Save(c context.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(j).Error
}

func (j *JobHistory) Find(c context.Context, tx *gorm.DB, search *JobHistory) (*JobHistory, error) {
	out := &JobHistory{}
	return out, tx.WithContext(c).Where(search).Find(&out).Error
}

func (j *JobHistory) Updates(c context.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Where("id = ?", j.ID).Updates(j).Error
}
