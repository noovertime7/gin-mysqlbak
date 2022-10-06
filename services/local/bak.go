package local

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/core"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/public/database"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/pkg/errors"
	"sync"
)

var (
	localBakService   *bakBakService
	bakBakServiceOnce sync.Once
)

// GetBakBakService 单例模式
func GetBakBakService() *bakBakService {
	bakBakServiceOnce.Do(func() {
		localBakService = &bakBakService{}
	})
	return localBakService
}

type bakBakService struct {
}

func (bak *bakBakService) Start(c *gin.Context, info *dto.Bak) error {
	tx := database.GetDB()
	taskInfo := &dao.TaskInfo{Id: info.ID}
	taskDetail, err := taskInfo.TaskDetail(c, tx, taskInfo)
	if err != nil {
		return err
	}
	bakHandler, err := core.NewBakHandler(taskDetail)
	if err != nil {
		return err
	}
	if err = bakHandler.StartBak(); err != nil {
		return err
	}
	//更改任务状态
	taskInfo.Status = 1
	if err = taskInfo.Updates(c, tx); err != nil {
		return err
	}
	return nil
}

func (bak *bakBakService) Stop(c *gin.Context, info *dto.Bak) error {
	var err error
	//将任务状态改为false
	taskinfo := &dao.TaskInfo{
		Id: info.ID,
	}
	taskinfo, err = taskinfo.Find(c, database.GetDB(), taskinfo)
	if err != nil {
		return err
	}
	taskinfo.Status = 0
	if err = taskinfo.UpdatesStatus(database.GetDB()); err != nil {
		return err
	}
	var bakHandler = &core.BakHandler{}
	if err = bakHandler.StopBak(info.ID); err != nil {
		return err
	}
	return nil
}

func (bak *bakBakService) StartTaskByHost(c *gin.Context, info *dto.HostIDInput) error {
	taskDB := &dao.TaskInfo{}
	result, err := taskDB.FindAllTask(c, database.GetDB(), info)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("当前主机备份任务为空")
	}
	for _, task := range result {
		if task.Status == 1 {
			log.Logger.Infof("TASK_ID:%d,HOSTID:%d,%s数据库任务已经打开,返回", task.Id, task.HostID, task.DBName)
			continue
		}
		taskDetail, err := taskDB.TaskDetail(c, database.GetDB(), task)
		if err != nil {
			return err
		}
		bakHandler, err := core.NewBakHandler(taskDetail)
		if err != nil {
			return err
		}
		if err = bakHandler.StartBak(); err != nil {
			return err
		}
		// 修改任务启动状态
		taskDB.Status = 1
		taskDB.Id = task.Id
		if err = taskDB.UpdatesStatus(database.GetDB()); err != nil {
			return err
		}
	}
	return nil
}

func (bak *bakBakService) StopTaskByHost(c *gin.Context, info *dto.HostIDInput) error {
	taskDB := &dao.TaskInfo{}
	result, err := taskDB.FindAllTask(c, database.GetDB(), info)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return err
	}
	for _, task := range result {
		if task.Status == 0 {
			log.Logger.Infof("TASK_ID:%d,HOSTID:%d,%s数据库任务已经关闭,返回", task.Id, task.HostID, task.DBName)
			continue
		}
		bakHandler := &core.BakHandler{DbName: task.DBName}
		err = bakHandler.StopBak(task.Id)
		if err != nil {
			return err
		}
		// 修改任务启动状态为关闭
		taskDB.Id = task.Id
		taskDB.Status = 0
		if err = taskDB.UpdatesStatus(database.GetDB()); err != nil {
			return err
		}
	}
	return nil
}
