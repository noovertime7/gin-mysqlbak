package controller

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
)

type DashboardController struct{}

func DashboardRegister(group *gin.RouterGroup) {
	dashboard := &DashboardController{}
	group.GET("/panel_group_data", dashboard.PanelGroupData)
}

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
		middleware.ResponseError(c, 2002, err)
		return
	}
	list, err := taskinfo.FindStatusUpTask(c, tx)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	runningNum := len(list)
	out := dto.PanelGroupDataOutPut{
		TaskNum:       taskNum,
		HistoryNum:    histryNum,
		RunningProNum: runningNum,
		StopProNum:    0,
	}
	middleware.ResponseSuccess(c, out)
}
