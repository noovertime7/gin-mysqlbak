package controller

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/core"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/pkg/errors"
)

type DashboardController struct{}

func DashboardRegister(group *gin.RouterGroup) {
	dashboard := &DashboardController{}
	group.GET("/panel_group_data", dashboard.PanelGroupData)
	group.GET("/pie_chart_data", dashboard.PieChartData)
}

var HostNameMap = make(map[int]string)

func (service *DashboardController) PanelGroupData(c *gin.Context) {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	taskinfo := &dao.TaskInfo{}
	_, taskNum, err := taskinfo.PageList(c, tx, &dto.TaskListInput{PageNo: 1, PageSize: 1})
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	histry := &dao.BakHistory{}
	_, histryNum, err := histry.PageList(c, tx, &dto.HistoryListInput{PageNo: 1, PageSize: 1})
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	hostdb := &dao.HostDatabase{}
	_, hostNum, err := hostdb.PageList(c, tx, &dto.HostListInput{PageNo: 1, PageSize: 1})
	if err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	_, err = taskinfo.FindAllTask(c, tx, nil)
	if err != nil {
		middleware.ResponseError(c, 2005, err)
		return
	}
	runningNum := len(core.CronJob)
	out := dto.PanelGroupDataOutPut{
		HostNum:       hostNum,
		TaskNum:       taskNum,
		HistoryNum:    histryNum,
		RunningProNum: runningNum,
	}
	middleware.ResponseSuccess(c, out)
}

func (service *DashboardController) PieChartData(c *gin.Context) {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	hostdb := &dao.HostDatabase{}
	hosts, err := hostdb.FindAllHost(tx)
	if err != nil {
		middleware.ResponseError(c, 2005, err)
		return
	}
	for _, host := range hosts {
		HostNameMap[host.Id] = host.Host
	}

	taskInfo := &dao.TaskInfo{}
	list, err := taskInfo.GroupByHost(c, tx)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	var legend []string
	for index, item := range list {
		name, ok := HostNameMap[item.HostID]
		if !ok {
			middleware.ResponseError(c, 2003, errors.New("host not found"))
			return
		}
		list[index].Name = name
		legend = append(legend, name)
	}
	out := &dto.DashServiceStatOutput{
		Legend: legend,
		Data:   list,
	}
	middleware.ResponseSuccess(c, out)
}
