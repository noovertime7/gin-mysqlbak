package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
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
	hostService := pkg.GetMicroService(params.ServiceName).(proto.HostService)
	data, err := hostService.AddHost(c, &proto.HostAddInput{
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
	hostService := pkg.GetMicroService(params.ServiceName).(proto.HostService)
	data, err := hostService.DeleteHost(c, &proto.HostDeleteInput{
		ID: int32(params.ID),
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
	hostService := pkg.GetMicroService(params.ServiceName).(proto.HostService)
	data, err := hostService.UpdateHost(c, &proto.HostUpdateInput{
		ID:       int32(params.ID),
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
	hostService := pkg.GetMicroService(params.ServiceName).(proto.HostService)
	out, err := hostService.GetHostList(c, &proto.HostListInput{
		Info:     params.Info,
		PageNo:   int32(params.PageNo),
		PageSize: int32(params.PageSize),
	})
	if err != nil {
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	middleware.ResponseSuccess(c, out)
}
