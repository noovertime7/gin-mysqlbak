package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/conf"
	"github.com/noovertime7/gin-mysqlbak/controller"
	"github.com/noovertime7/gin-mysqlbak/services"
	"log"
	"net/http"
	"time"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	gin.SetMode(conf.GetStringConf("base", "debug_mode"))
	r := InitRouter()
	HttpSrvHandler = &http.Server{
		Addr:           conf.GetStringConf("http", "addr"),
		Handler:        r,
		ReadTimeout:    time.Duration(conf.GetIntConf("http", "read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(conf.GetIntConf("http", "write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << uint(conf.GetIntConf("http", "max_header_bytes")),
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", conf.GetStringConf("http", "addr"))
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", conf.GetStringConf("http", "addr"), err)
		}
	}()
	//如果主机丢失开关打开，才会开启主机检测功能
	if conf.GetBoolConf("HostLostAlarms", "enable") {
		go controller.HostPortCheck()
	}
}

func HttpServerStop() {
	services.StopAllUpTask()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
