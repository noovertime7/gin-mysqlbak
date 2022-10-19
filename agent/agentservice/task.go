package agentservice

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/task"
	"sync"
)

var (
	clusterTaskService *TaskService
	TaskServiceOnce    sync.Once
)

// GetClusterTaskService  单例模式
func GetClusterTaskService() *TaskService {
	TaskServiceOnce.Do(func() {
		clusterTaskService = &TaskService{}
	})
	return clusterTaskService
}

type TaskService struct{}

func (t *TaskService) TaskAdd(ctx *gin.Context, params *agentdto.TaskAddInput) error {
	taskService, addr, err := pkg.GetTaskService(params.ServiceName)
	if err != nil {
		return err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	taskAddinput := &task.TaskAddInput{
		HostID:          params.HostID,
		DBName:          params.DBName,
		ServiceName:     params.ServiceName,
		BackupCycle:     params.BackupCycle,
		KeepNumber:      params.KeepNumber,
		IsAllDBBak:      params.IsAllDBBak,
		IsDingSend:      params.IsDingSend,
		DingAccessToken: params.DingAccessToken,
		DingSecret:      params.DingSecret,
		OssType:         params.OssType,
		IsOssSave:       params.IsOssSave,
		Endpoint:        params.Endpoint,
		OssAccess:       params.OssAccess,
		OssSecret:       params.OssSecret,
		BucketName:      params.BucketName,
		Directory:       params.Directory,
	}
	data, err := taskService.TaskAdd(ctx, taskAddinput, ops)
	if err != nil || !data.OK {
		return err
	}
	return nil
}

func (t *TaskService) TaskAutoAdd(ctx *gin.Context, params *agentdto.TaskAutoAddInput) error {
	taskService, addr, err := pkg.GetTaskService(params.ServiceName)
	if err != nil {
		return err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	taskAddinput := &task.TaskAutoCreateInPut{
		HostID:          params.HostID,
		BackupCycle:     params.BackupCycle,
		KeepNumber:      params.KeepNumber,
		IsAllDBBak:      params.IsAllDBBak,
		IsDingSend:      params.IsDingSend,
		DingAccessToken: params.DingAccessToken,
		DingSecret:      params.DingSecret,
		OssType:         params.OssType,
		IsOssSave:       params.IsOssSave,
		Endpoint:        params.Endpoint,
		OssAccess:       params.OssAccess,
		OssSecret:       params.OssSecret,
		BucketName:      params.BucketName,
		Directory:       params.Directory,
	}
	data, err := taskService.TaskAutoCreate(ctx, taskAddinput, ops)
	if err != nil || !data.OK {
		return err
	}
	return nil
}

func (t *TaskService) TaskList(ctx context.Context, params *agentdto.TaskListInput) (*task.TaskListOutPut, error) {
	taskService, addr, err := pkg.GetTaskService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	taskListInput := &task.TaskListInput{
		Info:     params.Info,
		PageNo:   params.PageNo,
		PageSize: params.PageSize,
		HostID:   params.HostId,
	}
	return taskService.TaskList(ctx, taskListInput, ops)
}

func (t *TaskService) GetTaskUnscopedList(ctx context.Context, params *agentdto.TaskListInput) (*task.TaskListOutPut, error) {
	taskService, addr, err := pkg.GetTaskService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	taskListInput := &task.TaskListInput{
		Info:     params.Info,
		PageNo:   params.PageNo,
		PageSize: params.PageSize,
		HostID:   params.HostId,
	}
	return taskService.GetTaskUnscopedList(ctx, taskListInput, ops)
}

func (t *TaskService) TaskDetail(ctx *gin.Context, params *agentdto.TaskIDInput) (*task.TaskDetailOutPut, error) {
	taskService, addr, err := pkg.GetTaskService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	taskDetailInput := &task.TaskIDInput{ID: params.ID}
	return taskService.TaskDetail(ctx, taskDetailInput, ops)
}

func (t *TaskService) TaskDelete(ctx *gin.Context, params *agentdto.TaskDeleteInput) error {
	taskService, addr, err := pkg.GetTaskService(params.ServiceName)
	if err != nil {
		return err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	taskDeleteInput := &task.TaskIDInput{ID: params.ID}
	data, err := taskService.TaskDelete(ctx, taskDeleteInput, ops)
	if err != nil || !data.OK {
		return err
	}
	return nil
}

func (t *TaskService) TaskDestroy(ctx *gin.Context, params *agentdto.TaskDeleteInput) error {
	taskService, addr, err := pkg.GetTaskService(params.ServiceName)
	if err != nil {
		return err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	taskDeleteInput := &task.TaskIDInput{ID: params.ID}
	data, err := taskService.TaskDestroy(ctx, taskDeleteInput, ops)
	if err != nil || !data.OK {
		return err
	}
	return nil
}

func (t *TaskService) TaskRestore(ctx *gin.Context, params *agentdto.TaskDeleteInput) error {
	taskService, addr, err := pkg.GetTaskService(params.ServiceName)
	if err != nil {
		return err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	taskDeleteInput := &task.TaskIDInput{ID: params.ID}
	data, err := taskService.RestoreTask(ctx, taskDeleteInput, ops)
	if err != nil || !data.OK {
		return err
	}
	return nil
}

func (t *TaskService) GetDateNumInfo(ctx context.Context, params *agentdto.GetDateNumInfoInput) (*task.DateNumInfoOut, error) {
	taskService, addr, err := pkg.GetTaskService(params.ServiceName)
	if err != nil {
		return nil, err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	taskDeleteInput := &task.DateNumInfoInput{Date: params.Date}
	data, err := taskService.GetDateNumInfo(ctx, taskDeleteInput, ops)
	return data, nil
}

func (t *TaskService) TaskUpdate(ctx *gin.Context, params *agentdto.TaskUpdateInput) error {
	taskService, addr, err := pkg.GetTaskService(params.ServiceName)
	if err != nil {
		return err
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	taskUpdateInput := &task.TaskUpdateInput{
		ID:              params.ID,
		HostID:          params.HostID,
		DBName:          params.DBName,
		ServiceName:     params.ServiceName,
		BackupCycle:     params.BackupCycle,
		KeepNumber:      params.KeepNumber,
		IsAllDBBak:      params.IsAllDBBak,
		IsDingSend:      params.IsDingSend,
		DingAccessToken: params.DingAccessToken,
		DingSecret:      params.DingSecret,
		OssType:         params.OssType,
		IsOssSave:       params.IsOssSave,
		Endpoint:        params.Endpoint,
		OssAccess:       params.OssAccess,
		OssSecret:       params.OssSecret,
		BucketName:      params.BucketName,
		Directory:       params.Directory,
	}
	data, err := taskService.TaskUpdate(ctx, taskUpdateInput, ops)
	if err != nil || !data.OK {
		return err
	}
	return nil
}
