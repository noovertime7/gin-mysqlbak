package main

import (
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/noovertime7/gin-mysqlbak/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := lib.InitModule("./conf/dev/", []string{"base", "mysql"})
	if err != nil {
		fmt.Println(err)
	}
	defer lib.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
