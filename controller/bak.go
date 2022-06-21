package controller

import (
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/core"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"gorm.io/gorm"
	"time"
)

type BakController struct {
	AfterBakChan chan *core.BakHandler
}

func BakRegister(group *gin.RouterGroup) {
	bak := &BakController{}
	group.POST("/start", bak.StartBak)
	group.POST("/stop", bak.StopBak)
}

func (b *BakController) StartBak(c *gin.Context) {
	params := &dto.Bak{}
	if err := params.BindValidParm(c); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	taskinfo := &dao.TaskInfo{Id: params.ID}
	taskdetail, err := taskinfo.TaskDetail(c, tx, taskinfo)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	b.AfterBakChan = make(chan *core.BakHandler, 10)
	tx, err = lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	go b.ListenAndSave(c, tx, b.AfterBakChan)
	bakhandler, err := core.NewBakController(taskdetail, b.AfterBakChan)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	if err = bakhandler.StartBak(); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "启动任务成功")
}

func (b *BakController) StopBak(c *gin.Context) {
	params := &dto.Bak{}
	if err := params.BindValidParm(c); err != nil {
		log.Logger.Error("BakHandleController 解析参数失败")
		middleware.ResponseError(c, 1007, err)
		return
	}
	var bakhandler = &core.BakHandler{}
	if err := bakhandler.StopBak(params.ID); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 1008, err)
		return
	}
	middleware.ResponseSuccess(c, "任务停止成功")
}

func (b *BakController) ListenAndSave(ctx *gin.Context, tx *gorm.DB, AfterBakChan chan *core.BakHandler) {
	log.Logger.Info("开始监听备份状态消息")
	fmt.Println("ListenAndSave", AfterBakChan)
	for {
		select {
		case afterbakhandler := <-AfterBakChan:
			bakhistory := &dao.BakHistory{
				TaskID:    afterbakhandler.TaskID,
				Host:      afterbakhandler.Host,
				DBName:    afterbakhandler.DbName,
				BakStatus: afterbakhandler.BakStatus,
				Msg:       afterbakhandler.BakMsg,
				FileSize:  afterbakhandler.FileSize,
				FileName:  afterbakhandler.FileName,
				BakTime:   time.Now(),
			}
			log.Logger.Info("接收到备份消息，数据入库")
			if err := bakhistory.Save(ctx, tx); err != nil {
				log.Logger.Error("保存备份历史到数据库失败", err)
				return
			}
		}
	}
}
