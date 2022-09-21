package local

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/conf"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/public/database"
	"sync"
)

var (
	localhostService *hostService
	hostServiceOnce  sync.Once
)

const (
	mysql = iota + 1
	elastic
)

// GetHostService 单例模式
func GetHostService() *hostService {
	hostServiceOnce.Do(func() {
		localhostService = &hostService{}
	})
	return localhostService
}

type hostService struct{}

// HostAdd 添加主机
func (h *hostService) HostAdd(ctx *gin.Context, info *dto.HostAddInput) error {
	tx := database.GetDB()
	host := &dao.HostDatabase{Host: info.Host, Password: info.Password, User: info.User, HostStatus: 1, Content: info.Content, Type: sql.NullInt32{Int32: info.Type, Valid: true}}
	switch info.Type {
	case mysql:
		host.Avatar = sql.NullString{String: conf.GetStringConf("app", "mysqlAvatar"), Valid: true}
	case elastic:
		host.Avatar = sql.NullString{String: conf.GetStringConf("app", "elasticAvatar"), Valid: true}
	default:
		host.Avatar = sql.NullString{String: conf.GetStringConf("app", "mysqlAvatar"), Valid: true}
	}
	if err := host.Save(ctx, tx); err != nil {
		return err
	}
	return nil
}

// HostDelete 删除主机
func (h *hostService) HostDelete(ctx *gin.Context, info *dto.HostDeleteInput) error {
	tx := database.GetDB()
	// 读取基本信息
	hostDB := &dao.HostDatabase{Id: info.ID}
	var err error
	host, err := hostDB.Find(ctx, tx, hostDB)
	if err != nil {
		return err
	}
	if host.Id == 0 {
		return fmt.Errorf("主机不存在,请检查id是否正确")
	}
	//删除标记修改为1
	host.IsDeleted = 1
	return host.Save(ctx, tx)
}

// HostUpdate 更新host
func (h *hostService) HostUpdate(ctx *gin.Context, info *dto.HostUpdateInput) error {
	tx := database.GetDB()
	host := &dao.HostDatabase{
		Id:         info.ID,
		Host:       info.Host,
		User:       info.User,
		Password:   info.Password,
		Content:    info.Content,
		HostStatus: 1,
		Type:       sql.NullInt32{Int32: info.Type, Valid: true},
	}
	switch info.Type {
	case mysql:
		host.Avatar = sql.NullString{String: conf.GetStringConf("app", "mysqlAvatar"), Valid: true}
	case elastic:
		host.Avatar = sql.NullString{String: conf.GetStringConf("app", "elasticAvatar"), Valid: true}
	default:
		host.Avatar = sql.NullString{String: conf.GetStringConf("app", "mysqlAvatar"), Valid: true}
	}
	return host.Save(ctx, tx)
}

// GetHostList 查询host
func (h *hostService) GetHostList(ctx *gin.Context, info *dto.HostListInput) (*dto.HostListOutput, error) {
	tx := database.GetDB()
	hostDB := &dao.HostDatabase{}
	list, total, err := hostDB.PageList(ctx, tx, info)
	if err != nil {
		return nil, err
	}
	var outList []dto.HostListOutItem
	for _, listIterm := range list {
		task := &dao.TaskInfo{}
		tasks, err := task.FindAllTask(ctx, tx, &dto.HostIDInput{HostID: listIterm.Id})
		if err != nil {
			return nil, err
		}
		outItem := dto.HostListOutItem{
			ID:         listIterm.Id,
			Host:       listIterm.Host,
			User:       listIterm.User,
			Password:   listIterm.Password,
			HostStatus: listIterm.HostStatus,
			Content:    listIterm.Content,
			TaskNum:    len(tasks),
			Avatar:     listIterm.Avatar.String,
			Type:       listIterm.Type.Int32,
		}
		outList = append(outList, outItem)
	}
	out := &dto.HostListOutput{
		Total: total,
		List:  outList,
	}
	return out, nil
}
