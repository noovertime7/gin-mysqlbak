package local

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdao"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/public/database"
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
		taskSum += date.TaskNum
		finishSum += date.FinishNum
		outItem := &dto.DateInfoOutItem{
			Date:      datetime,
			AgentNum:  date.AgentNum,
			TaskNum:   date.TaskNum,
			FinishNum: date.FinishNum,
		}
		out = append(out, outItem)
	}
	fmt.Println("TaskNum out[0] = ", out[0])
	fmt.Println("TaskNum out[len(out)-1 = ", out[len(out)-1])
	latestInfo := out[0]
	oldInfo := out[len(out)-1]
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
