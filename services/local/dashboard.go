package local

import (
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
