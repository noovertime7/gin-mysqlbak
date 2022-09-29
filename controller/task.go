package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public/database"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
	"github.com/noovertime7/gin-mysqlbak/services/local"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type TaskController struct {
	service *local.TaskService
}

func TaskRegister(group *gin.RouterGroup) {
	task := &TaskController{service: local.GetTaskService()}
	group.POST("/taskadd", task.TaskAdd)
	group.GET("/tasklist", task.TaskList)
	group.GET("/taskdetail", task.TaskDetail)
	group.DELETE("/taskdelete", task.TaskDelete)
	group.PUT("/taskupdate", task.TaskUpdate)
}

func (t *TaskController) TaskAdd(c *gin.Context) {
	params := &dto.TaskAddInput{}
	if err := params.BindValidParams(c); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	tx := database.GetDB()
	//开启事务
	tx = tx.Begin()
	taskinfo := &dao.TaskInfo{
		DBName:      params.DBName,
		BackupCycle: params.BackupCycle,
		KeepNumber:  params.KeepNumber,
		IsAllDBBak:  params.IsAllDBBak,
		HostID:      params.HostID,
		Status:      0,
		IsDelete:    0,
	}
	if err := taskinfo.Save(c, tx); err != nil {
		tx.Rollback()
		log.Logger.Error(err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.TaskAddError, err))
		return
	}
	ossdb := &dao.OssDatabase{
		TaskID:     taskinfo.Id,
		IsOssSave:  params.IsOssSave,
		OssType:    params.OssType,
		Endpoint:   params.Endpoint,
		OssAccess:  params.OssAccess,
		OssSecret:  params.OssSecret,
		BucketName: params.BucketName,
		Directory:  params.Directory,
	}
	if err := ossdb.Save(c, tx); err != nil {
		tx.Rollback()
		log.Logger.Error(err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.TaskAddError, err))
		return
	}
	dingdb := &dao.DingDatabase{
		TaskID:          taskinfo.Id,
		IsDingSend:      params.IsDingSend,
		DingAccessToken: params.DingAccessToken,
		DingSecret:      params.DingSecret,
	}
	if err := dingdb.Save(c, tx); err != nil {
		tx.Rollback()
		log.Logger.Error(err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.TaskAddError, err))
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(c, "添加任务成功")
}

func (s *TaskController) TaskDelete(ctx *gin.Context) {
	params := &dto.TaskDeleteInput{}
	if err := params.BindValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskDeleteError, err))
		return
	}
	tx := database.GetDB()
	// 读取基本信息
	taskinfo := &dao.TaskInfo{Id: params.ID}
	taskinfo, err := taskinfo.Find(ctx, tx, taskinfo)
	if err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskGetError, err))
		return
	}
	if taskinfo.Id == 0 {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskNodeFound, err))
		return
	}
	taskinfo.IsDelete = 1
	if err := taskinfo.Save(ctx, tx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskUpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

func (t *TaskController) TaskUpdate(c *gin.Context) {
	params := &dto.TaskUpdateInput{}
	if err := params.BindValidParams(c); err != nil {
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	tx := database.GetDB()
	tx = tx.Begin()
	taskinfo := &dao.TaskInfo{
		Id:          params.ID,
		HostID:      params.HostID,
		DBName:      params.DBName,
		BackupCycle: params.BackupCycle,
		KeepNumber:  params.KeepNumber,
		IsAllDBBak:  params.IsAllDBBak,
	}
	if err := taskinfo.Updates(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.TaskUpdateError, err))
		return
	}
	ding := &dao.DingDatabase{
		TaskID:          taskinfo.Id,
		DingSecret:      params.DingSecret,
		DingAccessToken: params.DingAccessToken,
		IsDingSend:      params.IsDingSend}
	if err := ding.Updates(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.TaskUpdateError, err))
		return
	}

	oss := &dao.OssDatabase{
		TaskID:     taskinfo.Id,
		IsOssSave:  params.IsOssSave,
		OssType:    params.OssType,
		Endpoint:   params.Endpoint,
		OssAccess:  params.OssAccess,
		OssSecret:  params.OssSecret,
		BucketName: params.BucketName,
		Directory:  params.Directory,
	}
	if err := oss.Updates(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.TaskUpdateError, err))
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(c, "任务更改成功")
}

func (t *TaskController) TaskList(c *gin.Context) {
	params := &dto.TaskListInput{}
	if err := params.BindValidParams(c); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	out, err := t.service.GetTaskList(c, params)
	if err != nil {
		log.Logger.Error("获取任务列表失败", err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.TaskGetError, err))
		return
	}
	middleware.ResponseSuccess(c, out)
}

func (s *TaskController) TaskDetail(ctx *gin.Context) {
	params := &dto.TaskDeleteInput{}
	if err := params.BindValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	detail, err := s.service.GetTaskDetail(ctx, params)
	if err != nil {
		log.Logger.Error("获取任务详情失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskGetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, detail)
}
