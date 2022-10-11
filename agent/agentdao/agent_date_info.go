package agentdao

import (
	"context"
	"gorm.io/gorm"
)

// AgentDateInfo 服务根据时间统计完成量
type AgentDateInfo struct {
	Id          int64  `gorm:"column:id;type:int(11);AUTO_INCREMENT;primary_key" json:"id"`
	TaskNum     int64  `gorm:"column:task_num;type:int(11);NOT NULL" json:"task_num"`
	FinishNum   int64  `gorm:"column:finish_num;type:int(11);NOT NULL" json:"finish_num"`
	AgentNum    int64  `gorm:"column:agent_num;type:int(11);NOT NULL" json:"agent_num"`
	CurrentTime string `gorm:"column:current_time;type:datetime;NOT NULL" json:"current_time"`
}

func (a *AgentDateInfo) TableName() string {
	return "agent_date_info"
}

func (a *AgentDateInfo) Save(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Save(a).Error
}

func (a *AgentDateInfo) Find(ctx context.Context, tx *gorm.DB, search *AgentDateInfo) (*AgentDateInfo, error) {
	out := &AgentDateInfo{}
	return out, tx.WithContext(ctx).Where(search).Find(out).Error
}
