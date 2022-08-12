package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/pkg"
	"github.com/noovertime7/gin-mysqlbak/agent/proto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
)

type AgentHostController struct {
}

func AgentHostRegister(group *gin.RouterGroup) {
	agenthost := &AgentHostController{}
	group.GET("/hostlist", agenthost.HostList)

}

func (t *AgentHostController) HostList(c *gin.Context) {
	params := &agentdto.HostListInput{}
	if err := params.BindValidParm(c); err != nil {
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	hostService := pkg.GetMicroService(params.ServiceName)
	out, err := hostService.(proto.HostService).GetHostList(c, &proto.HostListInput{
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
