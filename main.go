package main

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/noovertime7/gin-mysqlbak/router"
	"github.com/noovertime7/gin-mysqlbak/services"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := lib.InitModule("./conf/dev/", []string{"base", "mysql"})
	if err != nil {
		log.Fatal("加载配置文件失败", err)
	}
	defer lib.Destroy()
	router.HttpServerRun()
	// 执行初始化任务操作
	services.InitBak()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
