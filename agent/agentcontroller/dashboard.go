package agentcontroller

type DashBoardController struct{}

//func DashBoardRegister(group *gin.RouterGroup) {
//	dashboard := &DashBoardController{}
//	group.GET("/barchart", dashboard.GetBarChartData)
//}

//func (d *DashBoardController) GetBarChartData(ctx *gin.Context) {
//	serviceList := pkg.GetServiceList()
//	var num = 0
//	var name []string
//	var taskNum []int64
//	for _, service := range serviceList.AgentOutPutItem {
//		hostService := pkg.GetHostService(service.Name).(host.HostService)
//		out, err := hostService.GetHostList(ctx, &host.HostListInput{
//			Info:     "",
//			PageNo:   1,
//			PageSize: 999,
//		})
//		if err != nil {
//			middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
//			return
//		}
//		for _, hostItem := range out.ListItem {
//			num = num + int(hostItem.TaskNum)
//			fmt.Println(service.Name, hostItem.TaskNum)
//		}
//		name = append(name, service.Name)
//		taskNum = append(taskNum, int64(num))
//	}
//	out := agentdto.BarChartOutPut{ServiceName: name, TaskNum: taskNum}
//	middleware.ResponseSuccess(ctx, out)
//}
