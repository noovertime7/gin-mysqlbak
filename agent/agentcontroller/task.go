package agentcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/task"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type AgentTaskController struct{}

func AgentTaskRegister(group *gin.RouterGroup) {
	task := &AgentTaskController{}
	group.POST("/taskadd", task.TaskAdd)
	group.GET("/tasklist", task.TaskList)
	group.GET("/taskdetail", task.TaskDetail)
	group.DELETE("/taskdelete", task.TaskDelete)
	group.PUT("/taskupdate", task.TaskUpdate)
}

func (a *AgentTaskController) TaskAdd(ctx *gin.Context) {
	params := &agentdto.TaskAddInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	taskService, addr, err := pkg.GetTaskService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, 30001, err)
		return
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
		log.Logger.Error("agent添加主机失败", err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(ctx, data.Message)
	log.Logger.Info("agent添加主机成功")
}

func (a *AgentTaskController) TaskList(ctx *gin.Context) {
	params := &agentdto.TaskListInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	taskService, addr, err := pkg.GetTaskService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, 30001, err)
		return
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
	data, err := taskService.TaskList(ctx, taskListInput, ops)
	if err != nil {
		log.Logger.Error("agent查询任务列表失败", err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
	fmt.Println(data)
	log.Logger.Info("agent查询任务列表成功")
}

func (a *AgentTaskController) TaskDetail(ctx *gin.Context) {
	params := &agentdto.TaskDeleteInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	taskService, addr, err := pkg.GetTaskService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, 30001, err)
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	taskDetailInput := &task.TaskIDInput{ID: params.ID}
	data, err := taskService.TaskDetail(ctx, taskDetailInput, ops)
	if err != nil {
		log.Logger.Error("agent查询主机详情失败", err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
	log.Logger.Info("agent查询主机详情成功")
}

func (a *AgentTaskController) TaskDelete(ctx *gin.Context) {
	params := &agentdto.TaskDeleteInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	taskService, addr, err := pkg.GetTaskService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, 30001, err)
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	taskDeleteInput := &task.TaskIDInput{ID: params.ID}
	data, err := taskService.TaskDelete(ctx, taskDeleteInput, ops)
	if err != nil {
		log.Logger.Error("agent删除主机失败", err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
	log.Logger.Info("agent删除主机成功")
}

func (a *AgentTaskController) TaskUpdate(ctx *gin.Context) {
	params := &agentdto.TaskUpdateInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	fmt.Println("params", params)
	taskService, addr, err := pkg.GetTaskService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, 30001, err)
		return
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
	fmt.Println(taskUpdateInput)
	data, err := taskService.TaskUpdate(ctx, taskUpdateInput, ops)
	if err != nil || !data.OK {
		log.Logger.Error("agent更新主机失败", err)
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(ctx, data.Message)
	log.Logger.Info("agent更新主机成功")
}
