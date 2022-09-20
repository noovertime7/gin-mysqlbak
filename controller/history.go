package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/core"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/public/database"
	"github.com/noovertime7/gin-mysqlbak/services/local"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type HistoryController struct {
	AfterBakChan chan *core.BakHandler
}

func HistoryRegister(group *gin.RouterGroup) {
	bak := &HistoryController{}
	group.DELETE("/history_delete", bak.DeleteHistory)
	group.GET("/history_num_info", bak.GetHistoryNumInfo)
	group.GET("/findallhistory", bak.FindAllHistory)
	group.GET("/historylist", bak.HistoryList)
}

func (bak *HistoryController) DeleteHistory(ctx *gin.Context) {
	params := &dto.HistoryIDInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error("绑定参数失败")
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, fmt.Errorf("绑定参数失败"))
		return
	}
	s := local.GetBakHistoryService()
	if err := s.Delete(ctx, params.ID); err != nil {
		log.Logger.Error("删除历史记录失败")
		middleware.ResponseError(ctx, 2000, fmt.Errorf("删除历史记录失败"))
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

func (bak *HistoryController) FindAllHistory(c *gin.Context) {
	bakhis := &dao.BakHistory{}
	bakhistorys, err := bakhis.FindAllHistory(c, database.GetDB(), "")
	if err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 2002, err)
	}
	var bakhistorysOutputs []*dto.BakHistoryOutPut
	for _, bakhistory := range bakhistorys {
		his := &dto.BakHistoryOutPut{
			Host:    bakhistory.Host,
			DBName:  bakhistory.DBName,
			Message: bakhistory.Msg,
			Baktime: bakhistory.BakTime.Format("2006年01月02日15:04:01"),
		}
		bakhistorysOutputs = append(bakhistorysOutputs, his)
	}
	middleware.ResponseSuccess(c, bakhistorysOutputs)
}

func (bak *HistoryController) HistoryList(c *gin.Context) {
	params := &dto.HistoryListInput{}
	if err := params.BindValidParams(c); err != nil {
		log.Logger.Error("BakHandleController 解析参数失败")
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	s := local.GetBakHistoryService()
	out, err := s.GetHistoryList(c, params)
	if err != nil {
		log.Logger.Error("查询历史记录列表失败", err)
		middleware.ResponseError(c, 2006, err)
		return
	}
	middleware.ResponseSuccess(c, out)
}

func (bak *HistoryController) GetHistoryNumInfo(ctx *gin.Context) {
	s := local.GetBakHistoryService()
	out, err := s.GetHistoryNumInfo(ctx)
	if err != nil {
		log.Logger.Error("获取历史记录数量信息失败")
		middleware.ResponseError(ctx, 2007, err)
		return
	}
	middleware.ResponseSuccess(ctx, out)
}
