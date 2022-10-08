package agentservice

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdao"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/repository"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/public/database"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	clusterTaskOverViewService *TaskOverViewService
	TaskOverViewServiceOnce    sync.Once
)

// GetClusterTaskOverViewService  单例模式
func GetClusterTaskOverViewService() *TaskOverViewService {
	TaskOverViewServiceOnce.Do(func() {
		clusterTaskOverViewService = &TaskOverViewService{
			DB:            database.GetDB(),
			taskService:   GetClusterTaskService(),
			esTaskService: GetClusterEsTaskService(),
			esBakService:  GetClusterEsBakService(),
			bakService:    GetClusterBakService(),
		}
	})
	return clusterTaskOverViewService
}

type TaskOverViewService struct {
	DB            *gorm.DB
	taskService   *TaskService
	esTaskService *EsTaskService
	esBakService  *EsBakService
	bakService    *BakService
}

func (t *TaskOverViewService) Run(period time.Duration) error {
	for range time.Tick(period) {
		if err := t.Store(context.TODO()); err != nil {
			return err
		}
	}
	return nil
}

func (t *TaskOverViewService) Store(ctx context.Context) error {
	log.Logger.Info("启动集群全局任务管理定时器，开始循环同步任务数据到服务端")
	h := GetClusterHostService()
	ts := GetClusterTaskService()
	//1、查询当前所有服务
	var AgentService *repository.AgentService
	services, err := AgentService.GetAgentList(ctx, &agentdto.AgentListInput{PageNo: 1, PageSize: 9999})
	if err != nil {
		return err
	}
	for _, service := range services.AgentOutPutItem {
		//2、通过服务循环查询所有主机
		hosts, err := h.HostList(ctx, &agentdto.HostListInput{
			ServiceName: service.ServiceName,
			Info:        "",
			PageNo:      1,
			PageSize:    9999,
		})
		if err != nil {
			return err
		}
		for _, host := range hosts.ListItem {
			//3、通过主机循环查询主机下所有任务
			tasks, err := ts.GetTaskUnscopedList(ctx, &agentdto.TaskListInput{
				ServiceName: service.ServiceName,
				HostId:      host.ID,
				Info:        "",
				PageNo:      1,
				PageSize:    9999,
			})
			if err != nil {
				return err
			}
			for _, task := range tasks.TaskListItem {
				//4、到保存数据库
				taskOverDB := &agentdao.TaskOverview{
					ServiceName: service.ServiceName,
					HostId:      host.ID,
					Host:        host.Host,
					Type:        host.Type,
					TaskId:      task.ID,
					DbName:      task.DBName,
					BackupCycle: task.BackupCycle,
					KeepNumber:  task.KeepNumber,
					FinishNum:   task.FinishNum,
					Status:      sql.NullInt64{Int64: task.Status, Valid: true},
					CreatedAt:   public.StringToTime(task.CreateAt),
					UpdateAt:    public.StringToTime(task.UpdateAt),
					IsDeleted:   sql.NullInt64{Int64: task.IsDeleted, Valid: true},
					DeletedAt:   public.StringToTime(task.DeletedAt),
				}
				//先去查询一下看看有没有
				temp, err := taskOverDB.Find(ctx, t.DB, &agentdao.TaskOverview{ServiceName: service.ServiceName, TaskId: task.ID, HostId: host.ID})
				if err != nil {
					return err
				}
				if temp.ID != 0 {
					//更新操作
					log.Logger.Debug("任务总览定时器开始更新数据库", temp)
					if err := taskOverDB.Updates(ctx, t.DB, temp.ID); err != nil {
						return err
					}
				} else {
					//没有查到，新增操作
					log.Logger.Debug("任务总览定时器开始插入数据库", taskOverDB)
					if err := taskOverDB.Save(ctx, t.DB); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func (t *TaskOverViewService) GetTaskOverViewList(ctx *gin.Context, info *agentdto.TaskOverViewListInput) (*agentdto.TaskOverViewListOut, error) {
	taskDB := &agentdao.TaskOverview{}
	list, total, err := taskDB.PageList(ctx, t.DB, info)
	if err != nil {
		return nil, err
	}
	var out []*agentdto.TaskOverViewListOutItem
	for _, task := range list {
		outItem := &agentdto.TaskOverViewListOutItem{
			ID:          task.ID,
			ServiceName: task.ServiceName,
			HostID:      task.HostId,
			Host:        task.Host,
			TaskID:      task.TaskId,
			DBName:      task.DbName,
			BackupCycle: task.BackupCycle,
			KeepNumber:  task.KeepNumber,
			Status:      task.Status.Int64,
			Type:        task.Type,
			FinishNum:   task.FinishNum,
			IsDeleted:   task.IsDeleted.Int64,
		}
		out = append(out, outItem)
	}
	return &agentdto.TaskOverViewListOut{
		Total:    int64(total),
		List:     out,
		PageNo:   info.PageNo,
		PageSize: info.PageSize,
	}, nil
}

func (t *TaskOverViewService) StartTask(ctx *gin.Context, info *agentdto.StartOverViewBakInput) (string, error) {
	//switch info.Type {
	//case public.MysqlHost:
	data, err := t.bakService.StartBak(ctx, &agentdto.StartBakInput{
		TaskID:      info.TaskID,
		ServiceName: info.ServiceName,
	})
	if err != nil || !data.OK {
		return "", err
	}
	//启动成功，更新状态
	overviewDB := &agentdao.TaskOverview{ID: info.ID, Type: info.Type}
	return data.Message, t.UpdateStatus(ctx, overviewDB, 1)
	//case public.ElasticHost:
	//	data, err := t.esBakService.StartEsBak(ctx, &agentdto.ESBakStartInput{
	//		ID:          info.TaskID,
	//		ServiceName: info.ServiceName,
	//	})
	//	if err != nil || !data.OK {
	//		return "", err
	//	}
	//	//启动成功，更新状态
	//	overviewDB := &agentdao.TaskOverview{ID: info.ID, Type: info.Type}
	//	return data.Message, t.UpdateStatus(ctx, overviewDB, 1)
	//}
	//return "", errors.New("类型不匹配")
}

func (t *TaskOverViewService) StopTask(ctx *gin.Context, info *agentdto.StopOverViewBakInput) (string, error) {
	//switch info.Type {
	//case public.MysqlHost:
	data, err := t.bakService.StopBak(ctx, &agentdto.StopBakInput{
		TaskID:      info.TaskID,
		ServiceName: info.ServiceName,
	})
	if err != nil || !data.OK {
		return "", err
	}
	//启动成功，更新状态
	overviewDB := &agentdao.TaskOverview{ID: info.ID, Type: info.Type}
	return data.Message, t.UpdateStatus(ctx, overviewDB, 0)
	//case public.ElasticHost:
	//	data, err := t.esBakService.StopEsBak(ctx, &agentdto.ESBakStopInput{
	//		ID: info.TaskID,
	//	})
	//	if err != nil || !data.OK {
	//		return "", err
	//	}
	//	//启动成功，更新状态
	//	overviewDB := &agentdao.TaskOverview{ID: info.ID, Type: info.Type}
	//	return data.Message, t.UpdateStatus(ctx, overviewDB, 0)
	//}
	//return "", errors.New("类型不匹配")
}

func (t *TaskOverViewService) DeleteTask(ctx *gin.Context, info *agentdto.DeleteOverViewTaskInput) (string, error) {
	//switch info.Type {
	//case public.MysqlHost:
	if err := t.taskService.TaskDelete(ctx, &agentdto.TaskDeleteInput{ServiceName: info.ServiceName, ID: info.TaskID}); err != nil {
		return "", err
	}
	//启动成功，更新状态
	overviewDB := &agentdao.TaskOverview{ID: info.ID, Type: info.Type}
	if err := overviewDB.Delete(ctx, t.DB); err != nil {
		return "", err
	}
	return "删除成功", nil
	//case public.ElasticHost:
	//	data, err := t.esTaskService.DeleteEsTask(ctx, &agentdto.ESBakTaskIDInput{
	//		ID:          info.TaskID,
	//		ServiceName: info.ServiceName,
	//	})
	//	if err != nil || !data.OK {
	//		return "", err
	//	}
	//	//启动成功，更新状态
	//	overviewDB := &agentdao.TaskOverview{ID: info.ID, Type: info.Type}
	//	if err := overviewDB.Delete(ctx, t.DB); err != nil {
	//		return "", err
	//	}
	//	return data.Message, nil
	//}
	//return "", errors.New("类型不匹配")
}

func (t *TaskOverViewService) RestoreTask(ctx *gin.Context, info *agentdto.DeleteOverViewTaskInput) (string, error) {
	//switch info.Type {
	//case public.MysqlHost:
	if err := t.taskService.TaskRestore(ctx, &agentdto.TaskDeleteInput{
		ServiceName: info.ServiceName,
		ID:          info.TaskID,
	}); err != nil {
		return "", err
	}
	//还原成功，更新状态
	taskOveDB := &agentdao.TaskOverview{ID: info.ID, TaskId: info.TaskID}
	if err := t.UpdateDeleteStatus(ctx, taskOveDB, 0); err != nil {
		return "", err
	}
	return "还原成功", nil
	//case public.ElasticHost:
	//	data, err := t.esTaskService.RestoreEsTask(ctx, &agentdto.ESBakTaskIDInput{ID: info.TaskID, ServiceName: info.ServiceName})
	//	if err != nil || !data.OK {
	//		return "", err
	//	}
	//	//启动成功，更新状态
	//	taskOveDB := &agentdao.TaskOverview{ID: info.ID, TaskId: info.TaskID}
	//	if err := t.UpdateDeleteStatus(ctx, taskOveDB, 0); err != nil {
	//		return "", err
	//	}
	//	return "还原成功", nil
	//}
	//return "", errors.New("类型不匹配")
}

func (t *TaskOverViewService) UpdateDeleteStatus(ctx *gin.Context, info *agentdao.TaskOverview, status int64) error {
	taskOverView, err := info.Find(ctx, t.DB, info)
	if err != nil {
		return err
	}
	taskOverView.IsDeleted = sql.NullInt64{Int64: status, Valid: true}
	if err := taskOverView.Updates(ctx, t.DB, info.ID); err != nil {
		return err
	}
	return nil
}

func (t *TaskOverViewService) UpdateStatus(ctx *gin.Context, info *agentdao.TaskOverview, status int64) error {
	taskOverView, err := info.Find(ctx, t.DB, info)
	if err != nil {
		return err
	}
	taskOverView.Status = sql.NullInt64{Int64: status, Valid: true}
	if err := taskOverView.Updates(ctx, t.DB, info.ID); err != nil {
		return err
	}
	return nil
}
