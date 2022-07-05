package controller

import (
	"context"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/pkg/errors"
	"net"
	"time"
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
	if err := HostPingCheck(params); err != nil {
		middleware.ResponseError(c, 1111, errors.New("数据库连接失败，请检查IP地址或端口"))
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 10003, err)
		return
	}
	host := &dao.HostDatabase{Host: params.Host, Password: params.Password, User: params.User, HostStatus: 1}
	go HostPortCheck(host)
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
	// 更改主机后进行ping测试
	hostinput := &dto.HostAddInput{Host: params.Host, User: params.User, Password: params.Password}
	if err := HostPingCheck(hostinput); err != nil {
		middleware.ResponseError(c, 1111, err)
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
		task := &dao.TaskInfo{}
		_, taskNum, err := task.PageList(c, tx, &dto.TaskListInput{HostId: listIterm.Id, PageNo: 1, PageSize: 1})
		if err != nil {
			log.Logger.Error(err)
			middleware.ResponseError(c, 10004, err)
			return
		}
		outItem := dto.HostListOutItem{
			ID:         listIterm.Id,
			Host:       listIterm.Host,
			User:       listIterm.User,
			Password:   listIterm.Password,
			HostStatus: 0,
			TaskNum:    taskNum,
		}
		outList = append(outList, outItem)
	}
	out := &dto.HostListOutput{
		Total: total,
		List:  outList,
	}
	middleware.ResponseSuccess(c, out)
}

func HostPingCheck(host *dto.HostAddInput) error {
	en, err := xorm.NewEngine("mysql", host.User+":"+host.Password+"@tcp("+host.Host+")/mysql?charset=utf8&parseTime=true")
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

func HostPortCheck(host *dao.HostDatabase) {
	for {
		log.Logger.Infof("主机存活端口检测程序启动:HOST:%v", host.Host)
		timeout := time.Duration(5 * time.Second)
		_, err := net.DialTimeout("tcp", host.Host, timeout)
		tx, _ := lib.GetGormPool("default")
		hostdb := &dao.HostDatabase{Host: host.Host}
		if err != nil {
			log.Logger.Warnf("主机端口检测失败:HOST:%v,ERR:%s", host.Host, err.Error())
			hostdb.HostStatus = 0
			_ = hostdb.UpdatesStatus(tx)
		}
		time.Sleep(10 * time.Minute)
	}
}
