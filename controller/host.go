package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/noovertime7/gin-mysqlbak/conf"
	"github.com/noovertime7/gin-mysqlbak/core"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public/database"
	"github.com/noovertime7/gin-mysqlbak/public/ding"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
	"github.com/noovertime7/gin-mysqlbak/services/local"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/pkg/errors"
	"net"
	"time"
)

type HostController struct{}

var HostOnlineMap map[int]string

func HostRegister(group *gin.RouterGroup) {
	host := &HostController{}
	group.POST("/hostadd", host.HostAdd)
	group.DELETE("/hostdelete", host.HostDelete)
	group.POST("/hostupdate", host.HostUpdate)
	group.GET("/hostlist", host.HostList)
}

var hostService = local.GetHostService()

func (h *HostController) HostAdd(c *gin.Context) {
	params := &dto.HostAddInput{}
	if err := params.BindValidParams(c); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.HostAddError, err))
		return
	}
	if err := HostPingCheck(params.User, params.Password, params.Host); err != nil {
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.HostCheckError, errors.New("数据库连接失败，请检查IP地址或端口")))
		return
	}
	if err := hostService.HostAdd(c, params); err != nil {
		log.Logger.Error("添加host失败", err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.HostAddError, err))
		return
	}
	middleware.ResponseSuccess(c, "添加成功")
}

func (h *HostController) HostDelete(ctx *gin.Context) {
	params := &dto.HostDeleteInput{}
	if err := params.BindValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := hostService.HostDelete(ctx, params); err != nil {
		log.Logger.Error("删除主机失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.HistoryDeleteError, err))
		return
	}
	//从在线列表中删除主机
	delete(HostOnlineMap, params.ID)
	middleware.ResponseSuccess(ctx, "删除主机成功")
}

func (h *HostController) HostUpdate(c *gin.Context) {
	params := &dto.HostUpdateInput{}
	if err := params.BindValidParams(c); err != nil {
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	// 更改主机后进行ping测试
	if err := HostPingCheck(params.User, params.Password, params.Host); err != nil {
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.HostCheckError, err))
		return
	}
	if err := hostService.HostUpdate(c, params); err != nil {
		log.Logger.Error("更新失败", err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.HostUpdateError, err))
		return
	}
	middleware.ResponseSuccess(c, "修改主机成功")
}

func (h *HostController) HostList(c *gin.Context) {
	params := &dto.HostListInput{}
	if err := params.BindValidParams(c); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	out, err := hostService.GetHostList(c, params)
	if err != nil {
		log.Logger.Error("查询host列表失败")
		middleware.ResponseError(c, globalError.NewGlobalError(globalError.HostGetError, err))
		return
	}
	middleware.ResponseSuccess(c, out)
}

func HostPingCheck(user, password, host string) error {
	en, err := xorm.NewEngine("mysql", user+":"+password+"@tcp("+host+")/mysql?charset=utf8&parseTime=true")
	defer en.Close()
	if err != nil {
		log.Logger.Errorf("创建数据库连接失败:%s", err.Error())
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err = en.PingContext(ctx); err != nil {
		return err
	}
	return nil
}

func HostPortCheck() {
	HostOnlineMap = make(map[int]string)
	tx := database.GetDB()
	hostdb := &dao.HostDatabase{}
	for {
		hostlists, err := hostdb.FindAllHost(tx)
		if err != nil {
			log.Logger.Error(err)
			return
		}
		//添加host信息到onlinemap中
		for _, host := range hostlists {
			HostOnlineMap[host.Id] = host.Host
		}
		for _, host := range HostOnlineMap {
			//30秒运行一次监测任务
			HostPortCheckHandler(host)
		}
		time.Sleep(1 * time.Minute)
	}
}

func HostPortCheckHandler(host string) {
	tx := database.GetDB()
	hostdb := &dao.HostDatabase{Host: host}
	log.Logger.Infof("主机存活端口检测程序启动:HOST:%v", host)
	_, err := net.DialTimeout("tcp", host, 5*time.Second)
	if err != nil {
		log.Logger.Warnf("主机端口检测失败:HOST:%v,ERR:%s", host, err.Error())
		log.Logger.Infof("开始修改数据库在线状态:HOST:%v", host)
		hostdb.HostStatus = 0
		_ = hostdb.UpdatesStatus(tx)
		//根据是否开启钉钉代理，选择调用
		if conf.GetBoolConf("dingProxyAgent", "enable") {
			log.Logger.Infof("主机失联告警：使用钉钉代理发送告警信息")
			dingSender := core.NewDingSender(
				conf.GetStringConf("HostLostAlarms", "accessToken"),
				conf.GetStringConf("HostLostAlarms", "secret"),
				fmt.Sprintf("主机端口检测失败,请检查！:HOST:%v,ERR:%s", host, err.Error()))
			data, err := dingSender.SendMessage()
			if err != nil {
				log.Logger.Error("钉钉消息发送失败", err, data)
				return
			}
		} else {
			log.Logger.Infof("主机失联告警：使用自身能力发送告警信息")
			webhook := ding.Webhook{
				AtAll:       true,
				AccessToken: conf.GetStringConf("HostLostAlarms", "accessToken"),
				Secret:      conf.GetStringConf("HostLostAlarms", "secret"),
			}
			err = webhook.SendTextMessage(fmt.Sprintf("主机端口检测失败,请检查！:HOST:%v,ERR:%s", host, err.Error()))
			if err != nil {
				log.Logger.Error("钉钉消息发送失败", err)
				return
			}
		}
		log.Logger.Warnf("主机端口检测失败,发送钉钉告警，修改在线状态完成:HOST:%v,ERR:%s", host)
		return
	}
	log.Logger.Infof("主机存活端口检测成功:HOST:%v", host)
	hostdb.HostStatus = 1
	_ = hostdb.UpdatesStatus(tx)
}
