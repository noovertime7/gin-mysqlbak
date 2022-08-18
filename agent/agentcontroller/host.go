package agentcontroller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto/host"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"time"
)

type AgentHostController struct{}

func AgentHostRegister(group *gin.RouterGroup) {
	agenthost := &AgentHostController{}
	group.GET("/hostlist", agenthost.HostList)
	group.POST("/hostadd", agenthost.AddHost)
	group.DELETE("/hostdelete", agenthost.DeleteHost)
	group.PUT("/hostupdate", agenthost.UpdateHost)
}

func (a *AgentHostController) AddHost(c *gin.Context) {
	params := &agentdto.HostAddInput{}
	if err := params.BindValidParm(c); err != nil {
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	////进行主机检测避免添加无用信息
	//if err := HostPingCheck(params.User, params.Password, params.Host); err != nil {
	//	log.Logger.Error("agent添加主机检测失败", err)
	//	middleware.ResponseError(c, 20001, err)
	//	return
	//}
	hostService := pkg.GetHostService(params.ServiceName).(host.HostService)
	data, err := hostService.AddHost(c, &host.HostAddInput{
		Host:     params.Host,
		UserName: params.User,
		Password: params.Password,
		Content:  params.Content,
	})
	if err != nil || !data.OK {
		log.Logger.Error("agent添加主机失败", err)
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(c, data.Message)
	log.Logger.Info("agent添加主机成功")
}

func (a *AgentHostController) DeleteHost(c *gin.Context) {
	params := &agentdto.HostDeleteInput{}
	if err := params.BindValidParm(c); err != nil {
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	hostService := pkg.GetHostService(params.ServiceName).(host.HostService)
	data, err := hostService.DeleteHost(c, &host.HostDeleteInput{
		ID: int64(params.ID),
	})
	if err != nil || !data.OK {
		log.Logger.Error("agent删除主机失败", err)
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(c, data.Message)
	log.Logger.Info("agent删除主机成功")
}

func (a *AgentHostController) UpdateHost(c *gin.Context) {
	params := &agentdto.HostUpdateInput{}
	if err := params.BindValidParm(c); err != nil {
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	////进行主机检测避免添加无用信息
	//if err := HostPingCheck(params.User, params.Password, params.Host); err != nil {
	//	log.Logger.Error("agent添加主机检测失败", err)
	//	middleware.ResponseError(c, 20001, err)
	//	return
	//}
	hostService := pkg.GetHostService(params.ServiceName).(host.HostService)
	data, err := hostService.UpdateHost(c, &host.HostUpdateInput{
		ID:       int64(params.ID),
		Host:     params.Host,
		UserName: params.User,
		Password: params.Password,
		Content:  params.Content,
	})
	if err != nil || !data.OK {
		log.Logger.Error("agent更新主机失败", err)
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(c, data.Message)
	log.Logger.Info("agent更新主机成功")
}

func (a *AgentHostController) HostList(c *gin.Context) {
	params := &agentdto.HostListInput{}
	if err := params.BindValidParm(c); err != nil {
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	hostService := pkg.GetHostService(params.ServiceName).(host.HostService)
	out, err := hostService.GetHostList(c, &host.HostListInput{
		Info:     params.Info,
		PageNo:   int64(params.PageNo),
		PageSize: int64(params.PageSize),
	})
	if err != nil {
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(c, out)
}

func HostPingCheck(User, Password, Host string) error {
	en, err := xorm.NewEngine("mysql", User+":"+Password+"@tcp("+Host+")/mysql?charset=utf8&parseTime=true")
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
