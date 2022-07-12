package services

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
)

type AgentService struct{}

func (as *AgentService) AddAgent(c *gin.Context, input *dto.AgentRegisterInfo) error {
	tx, _ := lib.GetGormPool("default")
	agentdb := &dao.AgentDB{
		AgentName:   input.AgentName,
		AgentIpPort: input.AgentIpPort,
		IsDelete:    0,
	}
	if err := agentdb.Save(c, tx); err != nil {
		return err
	}
	return nil
}
