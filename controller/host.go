package controller

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/pkg/errors"
)

type HostController struct{}

func HostRegister(group *gin.RouterGroup) {
	host := &HostController{}
	group.POST("/hostadd", host.HostAdd)
	group.DELETE("/hostdelete", host.HostDelete)
	group.POST("/hostupdate", host.HostUpdate)
	group.GET("/hostlist", host.HostList)
}

func (h *HostController) HostAdd(c *gin.Context) {
	params := &dto.HostAddInput{}
	if err := params.BindValidParm(c); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 10003, err)
		return
	}
	host := &dao.HostDatabase{Host: params.Host, Password: params.Password, User: params.User}
	if err = host.Save(c, tx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 10004, err)
		return
	}
	middleware.ResponseSuccess(c, "添加Host成功")
}

func (s *HostController) HostDelete(ctx *gin.Context) {
	params := &dto.HostDeleteInput{}
	if err := params.BindValidParm(ctx); err != nil {
		middleware.ResponseError(ctx, 30001, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 30002, err)
		return
	}
	// 读取基本信息
	hostinfo := &dao.HostDatabase{Id: params.ID}
	hostinfo, err = hostinfo.Find(ctx, tx, hostinfo)
	if err != nil {
		middleware.ResponseError(ctx, 30003, err)
		return
	}
	if hostinfo.Id == 0 {
		middleware.ResponseError(ctx, 30003, errors.New("主机不存在,请检查id是否正确"))
		return
	}
	hostinfo.IsDeleted = 1
	if err = hostinfo.Save(ctx, tx); err != nil {
		middleware.ResponseError(ctx, 30004, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除主机成功")
}

func (h *HostController) HostUpdate(c *gin.Context) {
	params := &dto.HostUpdateInput{}
	if err := params.BindValidParm(c); err != nil {
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 30002, err)
		return
	}
	host := &dao.HostDatabase{
		Id:       params.ID,
		Host:     params.Host,
		User:     params.User,
		Password: params.Password,
	}
	if err = host.Save(c, tx); err != nil {
		middleware.ResponseError(c, 30003, err)
		return
	}
	middleware.ResponseSuccess(c, "修改主机成功")
}

func (t *HostController) HostList(c *gin.Context) {
	params := &dto.HostListInput{}
	if err := params.BindValidParm(c); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 10003, err)
		return
	}
	hostinfo := &dao.HostDatabase{}
	list, total, err := hostinfo.PageList(c, tx, params)
	if err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 10004, err)
		return
	}
	var outList []dto.HostListOutItem
	for _, listIterm := range list {
		outItem := dto.HostListOutItem{
			ID:         listIterm.Id,
			Host:       listIterm.Host,
			User:       listIterm.User,
			Password:   listIterm.Password,
			HostStatus: 0,
			TaskNum:    10,
		}
		outList = append(outList, outItem)
	}
	out := &dto.HostListOutput{
		Total: total,
		List:  outList,
	}
	middleware.ResponseSuccess(c, out)
}
