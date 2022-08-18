package agentservice

import (
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdao"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"time"
)

type AgentService struct{}

func (a *AgentService) Register(ctx *gin.Context, agentInfo *agentdto.AgentRegisterInput) error {
	agentDb := &agentdao.AgentDB{
		ServiceName: agentInfo.ServiceName,
		Content:     agentInfo.Content,
		Address:     agentInfo.Address,
		AgentStatus: 1,
		LastTime:    time.Now(),
		CreateAt:    time.Now(),
		IsDeleted:   0,
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		return err
	}
	agent, err := agentDb.Find(ctx, tx, &agentdao.AgentDB{ServiceName: agentInfo.ServiceName, IsDeleted: 0})
	if agent.Id != 0 {
		agentDb.Id = agent.Id
		agentDb.CreateAt = agent.CreateAt
		return agentDb.Update(ctx, tx)
	}
	return agentDb.Save(ctx, tx)
}

func (a *AgentService) DeRegister(ctx *gin.Context, serviceName string) error {
	agentDb := &agentdao.AgentDB{ServiceName: serviceName}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		return err
	}
	agent, err := agentDb.Find(ctx, tx, agentDb)
	if err != nil {
		return err
	}
	fmt.Println(agent, agentDb)
	agent.AgentStatus = 0
	return agent.UpdateStatus(ctx, tx)
}

func (a *AgentService) GetAgentList(ctx *gin.Context, agentInfo *agentdto.AgentListInput) (*agentdto.AgentListOutPut, error) {
	agentDB := &agentdao.AgentDB{}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		return nil, err
	}
	agents, total, err := agentDB.PageList(ctx, tx, agentInfo)
	if err != nil {
		return nil, err
	}
	var agentOutitems []*agentdto.AgentOutPutItem
	for _, agent := range agents {
		tempAgent := &agentdto.AgentOutPutItem{
			ServiceName: agent.ServiceName,
			Address:     agent.Address,
			Content:     agent.Content,
		}
		agentOutitems = append(agentOutitems, tempAgent)
	}
	out := &agentdto.AgentListOutPut{
		Total:           total,
		AgentOutPutItem: agentOutitems,
	}
	return out, nil
}
