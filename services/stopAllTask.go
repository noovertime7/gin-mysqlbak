package services

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/noovertime7/gin-mysqlbak/core"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

func StopAllUpTask() {
	log.Logger.Info("程序终止,停止所有备份任务")
	tx, err := lib.GetGormPool("default")
	if err != nil {
		log.Logger.Error(err)
		return
	}
	taskinfos, err := dao.FindAllStatusUpTask(tx)
	var bakhandler core.BakHandler
	for _, task := range taskinfos {
		err = bakhandler.StopBak(task.Id)
		if err != nil {
			log.Logger.Error(err)
			return
		}
		task.Status = 0
		err = task.UpdatesStatus(tx)
		if err != nil {
			log.Logger.Error(err)
			return
		}
	}
}
