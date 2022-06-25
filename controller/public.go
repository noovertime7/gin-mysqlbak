package controller

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/core"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type PublicController struct{}

func PublicRegister(group *gin.RouterGroup) {
	pb := &PublicController{}
	group.GET("/initbak", pb.InitBak)
}

// InitBak 公共接口无需鉴权，用于程序崩溃后，重启会自动启动数据库备份任务
func (p *PublicController) InitBak(ctx *gin.Context) {
	var b BakController
	tx, err := lib.GetGormPool("default")
	if err != nil {
		log.Logger.Error(err)
		return
	}
	taskinfo := &dao.TaskInfo{}
	result, err := taskinfo.FindStatusUpTask(ctx, tx)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	if len(result) == 0 {
		log.Logger.Info("当前备份列表都为停止状态,无需初始化")
		return
	}
	b.AfterBakChan = make(chan *core.BakHandler, 10)
	go b.ListenAndSave(ctx, tx, b.AfterBakChan)
	for _, task := range result {
		taskdetail, err := taskinfo.TaskDetail(ctx, tx, task)
		if err != nil {
			log.Logger.Error(err)
			return
		}
		bakhandler, err := core.NewBakController(taskdetail, b.AfterBakChan)
		if err != nil {
			log.Logger.Error(err)
			return
		}
		if err = bakhandler.StartBak(); err != nil {
			log.Logger.Error(err)
			return
		}
		taskinfo.Status = 1
		if err = taskinfo.Updates(ctx, tx); err != nil {
			log.Logger.Error(err)
			return
		}
		middleware.ResponseSuccess(ctx, "初始化任务成功")
	}
}
