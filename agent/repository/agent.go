package repository

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdao"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/public/database"
	"time"
)

type AgentService struct{}

func (a *AgentService) Register(ctx *gin.Context, agentInfo *agentdto.AgentRegisterInput) error {
	agentDb := &agentdao.AgentDB{
		ServiceName: agentInfo.ServiceName,
		Content:     agentInfo.Content,
		Address:     agentInfo.Address,
		AgentStatus: 1,
		TaskNum:     agentInfo.TaskNum,
		FinishNum:   agentInfo.FinishNum,
		LastTime:    time.Now(),
		CreateAt:    time.Now(),
		IsDeleted:   0,
	}
	agent, err := agentDb.Find(ctx, database.GetDB(), &agentdao.AgentDB{ServiceName: agentInfo.ServiceName, IsDeleted: 0})
	if err != nil {
		return err
	}
	if agent.Id != 0 {
		agentDb.Id = agent.Id
		agentDb.CreateAt = agent.CreateAt
		return agentDb.Update(ctx, database.GetDB())
	}
	return agentDb.Save(ctx, database.GetDB())
}

func (a *AgentService) DeRegister(ctx *gin.Context, serviceName string) error {
	agentDb := &agentdao.AgentDB{ServiceName: serviceName}
	agent, err := agentDb.Find(ctx, database.GetDB(), agentDb)
	if err != nil {
		return err
	}
	agent.AgentStatus = 0
	return agent.UpdateStatus(ctx, database.GetDB())
}

func (a *AgentService) GetAgentList(ctx context.Context, agentInfo *agentdto.AgentListInput) (*agentdto.AgentListOutPut, error) {
	agentDB := &agentdao.AgentDB{}
	agents, total, err := agentDB.PageList(ctx, database.GetDB(), agentInfo)
	if err != nil {
		return nil, err
	}
	var agentOutItems []*agentdto.AgentOutPutItem
	for _, agent := range agents {
		tempAgent := &agentdto.AgentOutPutItem{
			ServiceName: agent.ServiceName,
			Address:     agent.Address,
			Content:     agent.Content,
			LastTime:    agent.LastTime.Format("2006年01月02日15:04:01"),
			TaskNum:     agent.TaskNum,
			FinishNum:   agent.FinishNum,
			AgentStatus: agent.AgentStatus,
			CreateAt:    agent.CreateAt.Format("2006年01月02日15:04:01"),
		}
		agentOutItems = append(agentOutItems, tempAgent)
	}
	out := &agentdto.AgentListOutPut{
		Total:           total,
		AgentOutPutItem: agentOutItems,
	}
	return out, nil
}

func (a *AgentService) GetServiceAddr(ctx context.Context, serviceName string) (string, error) {
	agentDb := &agentdao.AgentDB{ServiceName: serviceName}
	agent, err := agentDb.Find(ctx, database.GetDB(), agentDb)
	if err != nil {
		return "unknown", err
	}
	return agent.Address, nil
}

func (a *AgentService) GetServiceNumInfo(ctx *gin.Context) (*agentdto.AgentServiceNumInfoOutput, error) {
	list, err := a.GetAgentList(ctx, &agentdto.AgentListInput{PageNo: 1, PageSize: 99999, Info: ""})
	if err != nil {
		return nil, err
	}
	var (
		services    int
		tasks       int
		finishTasks int
	)
	services = list.Total
	for _, s := range list.AgentOutPutItem {
		tasks += s.TaskNum
		finishTasks += s.FinishNum
	}
	return &agentdto.AgentServiceNumInfoOutput{
		AllServices:    services,
		AllTasks:       tasks,
		AllFinishTasks: finishTasks,
	}, nil
}
