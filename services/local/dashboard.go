package local

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdao"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/agentservice"
	"github.com/noovertime7/gin-mysqlbak/agent/repository"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/public/database"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"gorm.io/gorm"
	"sync"
)

var (
	localDashboardService *DashboardService
	DashboardServiceOnce  sync.Once
)

// GetDashboardService 单例模式
func GetDashboardService() *DashboardService {
	DashboardServiceOnce.Do(func() {
		localDashboardService = &DashboardService{DB: database.GetDB()}
	})
	return localDashboardService
}

type DashboardService struct {
	DB *gorm.DB
}

func (d *DashboardService) GetSvcTNum(ctx *gin.Context) ([]*dto.ServiceTaskOutPut, error) {
	agentDB := agentdao.AgentDB{}
	list, _, err := agentDB.PageList(ctx, d.DB, &agentdto.AgentListInput{
		Info:     "",
		PageNo:   1,
		PageSize: public.LargePageSize,
	})
	if err != nil {
		return nil, err
	}
	var out []*dto.ServiceTaskOutPut
	for _, s := range list {
		outItem := &dto.ServiceTaskOutPut{
			ServiceName: s.ServiceName,
			TaskNum:     int64(s.TaskNum),
		}
		out = append(out, outItem)
	}
	return out, nil
}

func (d *DashboardService) GetSvcFinishNum(ctx *gin.Context) (*dto.SvcFinishNumInfoOut, error) {
	//1、查询当前所有服务
	var AgentService *repository.AgentService
	services, err := AgentService.GetAgentList(ctx, &agentdto.AgentListInput{PageNo: 1, PageSize: 9999})
	if err != nil {
		return nil, err
	}
	var (
		allFinishTotal int64
		failTotal      int64
	)
	for _, s := range services.AgentOutPutItem {
		//获取history服务实例
		historyService := *agentservice.GetClusterHistoryService()
		historyNumInfo, err := historyService.GetHistoryNumInfo(ctx, &agentdto.HistoryServiceNameInput{
			ServiceName: s.ServiceName,
		})
		if err != nil {
			return nil, err
		}
		allFinishTotal += historyNumInfo.AllNums
		failTotal += historyNumInfo.FailNum
	}
	return &dto.SvcFinishNumInfoOut{
		AllFinishTotal: allFinishTotal,
		AllFailTotal:   failTotal,
	}, nil
}

func (d *DashboardService) GetTaskNumByDate(ctx *gin.Context, info *dto.AgentDateInfoInput) (*dto.DateInfoOut, error) {
	dates := public.GetBeforeDates(info.Day)
	var out []*dto.DateInfoOutItem
	var (
		taskSum           int64
		finishSum         int64
		taskIncreaseNum   int64
		taskDecreaseNum   int64
		finishIncreaseNum int64
		finishDecreaseNum int64
	)
	for _, datetime := range dates {
		dateDB := &agentdao.AgentDateInfo{CurrentTime: datetime}
		date, err := dateDB.Find(ctx, database.GetDB(), dateDB)
		if err != nil {
			return nil, err
		}
		transdate := public.DateTran(datetime)
		taskSum += date.TaskNum
		finishSum += date.FinishNum
		outItem := &dto.DateInfoOutItem{
			Date:      transdate,
			TaskNum:   date.TaskNum,
			FinishNum: date.FinishNum,
		}
		out = append(out, outItem)
	}
	log.Logger.Debug("oldInfo = ", out[0])
	log.Logger.Debug("latestInfo = ", out[len(out)-1])
	oldInfo := out[0]
	latestInfo := out[len(out)-1]
	//如果最新的小于之前的任务数，代表减少了
	if latestInfo.TaskNum < oldInfo.TaskNum {
		taskIncreaseNum = 0
		taskDecreaseNum = oldInfo.TaskNum - latestInfo.TaskNum
	} else {
		taskDecreaseNum = 0
		//代表增加了，用最新的任务数减去8天前的任务数，得到增加的任务数
		taskIncreaseNum = latestInfo.TaskNum - oldInfo.TaskNum
	}
	//如果数组的第一个任务数小于最后一个任务数，代表任务增加了
	if latestInfo.FinishNum < oldInfo.FinishNum {
		finishIncreaseNum = 0
		finishDecreaseNum = oldInfo.FinishNum - latestInfo.FinishNum
	} else {
		finishDecreaseNum = 0
		finishIncreaseNum = latestInfo.FinishNum - oldInfo.FinishNum
	}
	return &dto.DateInfoOut{
		TaskTotal:         taskSum,
		FinishTotal:       finishSum,
		TaskIncreaseNum:   taskIncreaseNum,
		TaskDecreaseNum:   taskDecreaseNum,
		FinishIncreaseNum: finishIncreaseNum,
		FinishDecreaseNum: finishDecreaseNum,
		List:              out,
	}, nil
}
