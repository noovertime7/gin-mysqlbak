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
		clusterTaskOverViewService = &TaskOverViewService{DB: database.GetDB()}
	})
	return clusterTaskOverViewService
}

type TaskOverViewService struct {
	DB *gorm.DB
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
