package agentservice

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/esbak"
	"sync"
)

var (
	clusterEsTaskService *EsTaskService
	EsTaskServiceOnce    sync.Once
)

// GetClusterEsTaskService  单例模式
func GetClusterEsTaskService() *EsTaskService {
	EsTaskServiceOnce.Do(func() {
		clusterEsTaskService = &EsTaskService{}
	})
	return clusterEsTaskService
}

type EsTaskService struct{}

func (e *EsTaskService) AddEsTask(ctx *gin.Context, params *agentdto.ESBakTaskAddInput) (*esbak.EsOneMessage, error) {
	EsTaskService, addr, err := pkg.GetESTaskService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	return EsTaskService.TaskAdd(ctx, &esbak.EsBakTaskADDInput{
		HostID:      params.HostID,
		BackupCycle: params.BackupCycle,
		KeepNumber:  params.KeepNumber,
	}, ops)
}

func (e *EsTaskService) DeleteEsTask(ctx *gin.Context, params *agentdto.ESBakTaskIDInput) (*esbak.EsOneMessage, error) {
	EsTaskService, addr, err := pkg.GetESTaskService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	return EsTaskService.TaskDelete(ctx, &esbak.EsTaskIDInput{ID: params.ID}, ops)
}

func (e *EsTaskService) RestoreEsTask(ctx *gin.Context, params *agentdto.ESBakTaskIDInput) (*esbak.EsOneMessage, error) {
	EsTaskService, addr, err := pkg.GetESTaskService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	return EsTaskService.TaskRestore(ctx, &esbak.EsTaskIDInput{ID: params.ID}, ops)
}

func (e *EsTaskService) UpdateEsTask(ctx *gin.Context, params *agentdto.ESBakTaskUpdateInput) (*esbak.EsOneMessage, error) {
	EsTaskService, addr, err := pkg.GetESTaskService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	return EsTaskService.TaskUpdate(ctx, &esbak.EsBakTaskUpdateInput{
		ID:          params.ID,
		HostID:      params.HostID,
		BackupCycle: params.BackupCycle,
		KeepNumber:  params.KeepNumber,
	}, ops)
}

func (e *EsTaskService) GetEsTaskList(ctx *gin.Context, params *agentdto.ESBakTaskListInput) (*esbak.EsTaskListOutPut, error) {
	EsTaskService, addr, err := pkg.GetESTaskService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	return EsTaskService.GetTaskList(ctx, &esbak.EsTaskListInput{
		Info:     params.Info,
		PageNo:   params.PageNo,
		PageSize: params.PageSize,
	}, ops)
}

func (e *EsTaskService) GetEsTaskDetail(ctx *gin.Context, params *agentdto.ESBakTaskIDInput) (*esbak.EsTaskDetailOutPut, error) {
	EsTaskService, addr, err := pkg.GetESTaskService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	return EsTaskService.GetTaskDetail(ctx, &esbak.EsTaskIDInput{ID: params.ID}, ops)
}
