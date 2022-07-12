package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AgentDB struct {
	Id          int    `json:"id" gorm:"primary_key" description:"自增主键"`
	AgentName   string `json:"agent_name"  gorm:"column:agent_name" description:"agent name"`
	AgentIpPort string `json:"agent_ip_port"  gorm:"column:agent_ip_port" description:"agent 地址信息"`
	IsDelete    int    `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (a *AgentDB) TableName() string {
	return "t_agent"
}

func (a *AgentDB) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(a).Error
}

func (a *AgentDB) Find(c *gin.Context, tx *gorm.DB, search *AgentDB) (*AgentDB, error) {
	out := &AgentDB{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}
