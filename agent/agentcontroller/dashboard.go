package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
)

type DashBoardController struct{}

func DashBoardRegister(group *gin.RouterGroup) {
	dashboard := &DashBoardController{}
	group.GET("/barchart", dashboard.GetBarChartData)
}

func (d *DashBoardController) GetBarChartData(ctx *gin.Context) {
	serviceList, _ := AgentService.GetAgentList(ctx, &agentdto.AgentListInput{
		Info:     "",
		PageNo:   1,
		PageSize: 999,
	})
	var name []string
	var taskNum []int64
	for _, s := range serviceList.AgentOutPutItem {
		name = append(name, s.ServiceName)
		taskNum = append(taskNum, int64(s.TaskNum))
	}
	out := &agentdto.BarChartOutPut{
		ServiceName: name,
		TaskNum:     taskNum,
	}
	middleware.ResponseSuccess(ctx, out)
}
