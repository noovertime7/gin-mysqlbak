package agentcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/host"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
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
	var num = 0
	var name []string
	var taskNum []int64
	for _, service := range serviceList.AgentOutPutItem {
		hostService, addr, err := pkg.GetHostService(service.ServiceName)
		if err != nil {
			log.Logger.Error("获取Agent地址失败", err)
			middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
			return
		}
		var ops client.CallOption = func(options *client.CallOptions) {
			options.Address = []string{addr}
		}
		out, err := hostService.GetHostList(ctx, &host.HostListInput{
			Info:     "",
			PageNo:   1,
			PageSize: 999,
		}, ops)
		if err != nil {
			middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
			return
		}
		for _, hostItem := range out.ListItem {
			num = num + int(hostItem.TaskNum)
			fmt.Println(service.ServiceName, hostItem.TaskNum)
		}
		name = append(name, service.ServiceName)
		taskNum = append(taskNum, int64(num))
	}
	out := agentdto.BarChartOutPut{ServiceName: name, TaskNum: taskNum}
	middleware.ResponseSuccess(ctx, out)
}
