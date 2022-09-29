package local

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/public/database"
	"github.com/pkg/errors"
	"sync"
)

var (
	localTaskService *TaskService
	TaskServiceOnce  sync.Once
)

// GetTaskService 单例模式
func GetTaskService() *TaskService {
	TaskServiceOnce.Do(func() {
		localTaskService = &TaskService{}
	})
	return localTaskService
}

type TaskService struct{}

// GetTaskList 获取任务列表
func (t *TaskService) GetTaskList(ctx *gin.Context, info *dto.TaskListInput) (*dto.TaskListOutput, error) {
	return nil, errors.New("测试测试")
	tx := database.GetDB()
	taskDB := &dao.TaskInfo{}
	list, total, err := taskDB.PageList(ctx, tx, info)
	if err != nil {
		return nil, err

	}
	var outList []dto.TaskListOutItem
	for _, listIterm := range list {
		//将corn表达式转换为室间隔是
		nexttime, _ := public.Cronexpr(listIterm.BackupCycle)
		cronstr := nexttime
		hostDB := &dao.HostDatabase{Id: listIterm.HostID}
		host, err := hostDB.Find(ctx, tx, hostDB)
		if err != nil {
			return nil, err
		}
		outItem := dto.TaskListOutItem{
			ID:          listIterm.Id,
			Host:        host.Host,
			HostID:      listIterm.HostID,
			DBName:      listIterm.DBName,
			BackupCycle: cronstr,
			KeepNumber:  listIterm.KeepNumber,
			Status:      listIterm.Status,
			CreateAt:    listIterm.CreatedAt.Format("2006年01月02日15:04"),
		}
		outList = append(outList, outItem)
	}
	out := &dto.TaskListOutput{
		Total:    total,
		List:     outList,
		PageSize: info.PageSize,
		PageNo:   info.PageNo,
	}
	return out, nil
}

// GetTaskDetail 获取任务详情
func (t *TaskService) GetTaskDetail(ctx *gin.Context, info *dto.TaskDeleteInput) (*dto.TaskDeTailOutPut, error) {
	tx := database.GetDB()
	// 读取基本信息
	taskDB := &dao.TaskInfo{Id: info.ID}
	task, err := taskDB.Find(ctx, tx, taskDB)
	if err != nil {
		return nil, err
	}
	if task.Id == 0 {
		return nil, errors.New("任务不存在，请检查ID是否正确")
	}
	taskDetail, err := task.TaskDetail(ctx, tx, task)
	if err != nil {
		return nil, err
	}
	out := &dto.TaskDeTailOutPut{
		ID:              taskDetail.Info.Id,
		Host:            taskDetail.Host.Host,
		HostID:          taskDetail.Host.Id,
		DBName:          taskDetail.Host.Host,
		BackupCycle:     taskDetail.Info.BackupCycle,
		KeepNumber:      taskDetail.Info.KeepNumber,
		CreateAt:        taskDetail.Info.CreatedAt.Format("2006年01月02日15:04:01"),
		DingID:          taskDetail.Ding.Id,
		IsDingSend:      taskDetail.Ding.IsDingSend,
		DingAccessToken: taskDetail.Ding.DingAccessToken,
		DingSecret:      taskDetail.Ding.DingSecret,
		OssID:           taskDetail.Oss.Id,
		IsOssSave:       taskDetail.Oss.IsOssSave,
		OssType:         taskDetail.Oss.OssType,
		Endpoint:        taskDetail.Oss.Endpoint,
		OssAccess:       taskDetail.Oss.OssAccess,
		OssSecret:       taskDetail.Oss.OssSecret,
		BucketName:      taskDetail.Oss.BucketName,
		Directory:       taskDetail.Oss.Directory,
	}
	return out, nil
}
