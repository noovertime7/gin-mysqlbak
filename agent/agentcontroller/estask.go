package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/esbak"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type EsTaskController struct{}

func EsTaskRegister(group *gin.RouterGroup) {
	esTask := &EsTaskController{}
	group.POST("/taskadd", esTask.AddEsTask)
	group.DELETE("/taskdelete", esTask.DeleteEsTask)
	group.PUT("/taskupdate", esTask.UpdateEsTask)
	group.GET("/tasklist", esTask.GetEsTaskList)
	group.GET("/taskdetail", esTask.GetEsTaskDetail)
}

func (e *EsTaskController) AddEsTask(ctx *gin.Context) {
	params := &agentdto.ESBakTaskAddInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	EsTaskService, addr, err := pkg.GetESTaskService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.AgentGetAddressError, err))
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	data, err := EsTaskService.TaskAdd(ctx, &esbak.EsBakTaskADDInput{
		EsHost:      params.Host,
		EsUser:      params.User,
		EsPassword:  params.Password,
		BackupCycle: params.BackupCycle,
		KeepNumber:  params.KeepNumber,
	}, ops)
	if err != nil || !data.OK {
		log.Logger.Error("es添加任务失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskAddError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data.Message)
	log.Logger.Info("es添加任务成功", data.Message)
}

func (e *EsTaskController) DeleteEsTask(ctx *gin.Context) {
	params := &agentdto.ESBakTaskIDInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	EsTaskService, addr, err := pkg.GetESTaskService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.AgentGetAddressError, err))
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	data, err := EsTaskService.TaskDelete(ctx, &esbak.EsTaskIDInput{ID: params.ID}, ops)
	if err != nil || !data.OK {
		log.Logger.Error("es删除任务失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskDeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data.Message)
	log.Logger.Info("es删除任务成功", data.Message)
}

func (e *EsTaskController) UpdateEsTask(ctx *gin.Context) {
	params := &agentdto.ESBakTaskUpdateInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	EsTaskService, addr, err := pkg.GetESTaskService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.AgentGetAddressError, err))
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	data, err := EsTaskService.TaskUpdate(ctx, &esbak.EsBakTaskUpdateInput{
		ID:          params.ID,
		EsHost:      params.Host,
		EsUser:      params.User,
		EsPassword:  params.Password,
		BackupCycle: params.BackupCycle,
		KeepNumber:  params.KeepNumber,
	}, ops)
	if err != nil || !data.OK {
		log.Logger.Error("es修改任务失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskUpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data.Message)
	log.Logger.Info("es修改任务成功", data.Message)
}

func (e *EsTaskController) GetEsTaskList(ctx *gin.Context) {
	params := &agentdto.ESBakTaskListInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	EsTaskService, addr, err := pkg.GetESTaskService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.AgentGetAddressError, err))
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	data, err := EsTaskService.GetTaskList(ctx, &esbak.EsTaskListInput{
		Info:     params.Info,
		PageNo:   params.PageNo,
		PageSize: params.PageSize,
	}, ops)
	if err != nil {
		log.Logger.Error("获取es_task列表失败")
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskGetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
func (e *EsTaskController) GetEsTaskDetail(ctx *gin.Context) {
	params := &agentdto.ESBakTaskIDInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	EsTaskService, addr, err := pkg.GetESTaskService(params.ServiceName)
	if err != nil {
		log.Logger.Error("获取Agent地址失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.AgentGetAddressError, err))
		return
	}
	var ops client.CallOption = func(options *client.CallOptions) {
		options.Address = []string{addr}
	}
	data, err := EsTaskService.GetTaskDetail(ctx, &esbak.EsTaskIDInput{ID: params.ID}, ops)
	if err != nil {
		log.Logger.Error("获取Agent详情失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskGetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
