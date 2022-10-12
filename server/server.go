package server

import (
	"github.com/noovertime7/gin-mysqlbak/conf"
	"github.com/noovertime7/gin-mysqlbak/core/job"
	"github.com/noovertime7/gin-mysqlbak/job/system"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/services"
	"github.com/noovertime7/mysqlbak/pkg/log"
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
	ctx := GetGlobalContext()
	//默认同步时间周期为30分钟一次
	period := conf.GetStringConf("cluster", "clusterSyncPeriod")
	if period == "" {
		period = "00 30 * * *"
	}
	factory := job.GetJobFactory()
	taskSync := system.NewTaskSyncJob(ctx, period)
	//log.Logger.Infof("启动集群任务同步定时器，当前同步周期%s", period)
	datesync := system.NewDateNumInfoSync(ctx, period)
	factory.Register(datesync, job.JobType(public.DateNumInfoJob))
	factory.Register(taskSync, job.JobType(public.TaskSyncJob))
	return factory.Start()
}
