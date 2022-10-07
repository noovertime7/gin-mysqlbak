package server

import (
	"github.com/noovertime7/gin-mysqlbak/agent/agentservice"
	"github.com/noovertime7/gin-mysqlbak/conf"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/services"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"time"
)

// Start 启动服务有关
func Start() {
	//打印logo
	public.PrintLogo()
	go func() {
		if err := startSyncClusterTask(); err != nil {
			log.Logger.Warning("集群同步任务列表失败，会导致任务总览显示异常", err)
		}
	}()
}

func Stop() {
	services.StopAllUpTask()
}

// startSyncClusterTask 同步集群任务
func startSyncClusterTask() error {
	//默认同步时间周期为10分钟一次
	var defaultPeriod = 10
	s := agentservice.GetClusterTaskOverViewService()
	period := conf.GetIntConf("cluster", "clusterSyncPeriod")
	if period == 0 {
		period = defaultPeriod
	}
	log.Logger.Infof("启动集群任务同步定时器，当前同步周期%d/min", period)
	return s.Run(time.Duration(period) * time.Minute)
}
