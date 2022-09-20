package services

import (
	"github.com/noovertime7/gin-mysqlbak/core"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/public/database"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

func StopAllUpTask() {
	log.Logger.Info("程序终止,停止所有备份任务")
	tx := database.GetDB()
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
