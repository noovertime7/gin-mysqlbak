package controller

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/noovertime7/gin-mysqlbakv2/dao"
	"github.com/noovertime7/gin-mysqlbakv2/dto"
	"github.com/noovertime7/gin-mysqlbakv2/middleware"
	"github.com/noovertime7/gin-mysqlbakv2/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/pkg/errors"
)

type TaskController struct{}

func TaskRegister(group *gin.RouterGroup) {
	task := &TaskController{}
	group.POST("/taskadd", task.TaskAdd)
	group.POST("/tasklist", task.TaskList)
	group.POST("/taskdetail", task.TaskDetail)
	group.DELETE("/taskdelete", task.TaskDelete)
	group.PUT("/taskupdate", task.TaskUpdate)
}

func (t *TaskController) TaskAdd(c *gin.Context) {
	params := &dto.TaskAddInput{}
	if err := params.BindValidParm(c); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	if err := TaskPingCheck(params); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 10000, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		log.Logger.Panic(err)
	}
	//开启事务
	tx = tx.Begin()
	taskinfo := &dao.TaskInfo{
		Host:        params.Host,
		Password:    params.Password,
		User:        params.User,
		DBName:      params.DBName,
		BackupCycle: params.BackupCycle,
		KeepNumber:  params.KeepNumber,
		IsAllDBBak:  params.IsAllDBBak,
		IsDelete:    0,
	}
	if err = taskinfo.Save(c, tx); err != nil {
		tx.Rollback()
		log.Logger.Error(err)
		middleware.ResponseError(c, 10001, err)
		return
	}
	ossdb := &dao.OssDatabase{
		TaskID:     taskinfo.Id,
		IsOssSave:  params.IsOssSave,
		OssType:    params.OssType,
		Endpoint:   params.Endpoint,
		OssAccess:  params.OssAccess,
		OssSecret:  params.OssSecret,
		BucketName: params.BucketName,
		Directory:  params.Directory,
	}
	if err := ossdb.Save(c, tx); err != nil {
		tx.Rollback()
		log.Logger.Error(err)
		middleware.ResponseError(c, 10002, err)
		return
	}
	dingdb := &dao.DingDatabase{
		TaskID:          taskinfo.Id,
		IsDingSend:      params.IsDingSend,
		DingAccessToken: params.DingAccessToken,
		DingSecret:      params.DingSecret,
	}
	if err := dingdb.Save(c, tx); err != nil {
		tx.Rollback()
		log.Logger.Error(err)
		middleware.ResponseError(c, 10002, err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(c, "添加任务成功")
}

func (s *TaskController) TaskDelete(ctx *gin.Context) {
	params := &dto.TaskDeleteInput{}
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
	taskinfo := &dao.TaskInfo{Id: params.ID}
	taskinfo, err = taskinfo.Find(ctx, tx, taskinfo)
	if err != nil {
		middleware.ResponseError(ctx, 30003, err)
		return
	}
	if taskinfo.Id == 0 {
		middleware.ResponseError(ctx, 30003, errors.New("任务不存在,请检查id是否正确"))
		return
	}
	taskinfo.IsDelete = 1
	if err := taskinfo.Save(ctx, tx); err != nil {
		middleware.ResponseError(ctx, 30004, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

func (t *TaskController) TaskList(c *gin.Context) {
	params := &dto.TaskListInput{}
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
	taskinfo := &dao.TaskInfo{}
	list, total, err := taskinfo.PageList(c, tx, params)
	if err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(c, 10004, err)
		return
	}
	var outList []dto.TaskListOutItem
	for _, listIterm := range list {
		outItem := dto.TaskListOutItem{
			ID:          listIterm.Id,
			Host:        listIterm.Host,
			DBName:      listIterm.DBName,
			BackupCycle: listIterm.BackupCycle,
			KeepNumber:  listIterm.KeepNumber,
			CreateAt:    listIterm.CreatedAt,
		}
		outList = append(outList, outItem)
	}
	out := &dto.TaskListOutput{
		Total: total,
		List:  outList,
	}
	middleware.ResponseSuccess(c, out)
}

func (s *TaskController) TaskDetail(ctx *gin.Context) {
	params := &dto.TaskDeleteInput{}
	if err := params.BindValidParm(ctx); err != nil {
		middleware.ResponseError(ctx, public.ParamsBindErrorCode, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 30002, err)
		return
	}
	// 读取基本信息
	taskinfo := &dao.TaskInfo{Id: params.ID}
	taskinfo, err = taskinfo.Find(ctx, tx, taskinfo)
	if err != nil {
		middleware.ResponseError(ctx, 30003, err)
		return
	}
	if taskinfo.Id == 0 {
		middleware.ResponseError(ctx, 30004, errors.New("任务不存在,请检查id是否正确"))
		return
	}
	taskdetail, err := taskinfo.TaskDetail(ctx, tx, taskinfo)
	if err != nil {
		middleware.ResponseError(ctx, 30004, err)
		return
	}
	middleware.ResponseSuccess(ctx, taskdetail)
}

func (t *TaskController) TaskUpdate(c *gin.Context) {
	params := &dto.TaskUpdateInput{}
	if err := params.BindValidParm(c); err != nil {
		middleware.ResponseError(c, public.ParamsBindErrorCode, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 30002, err)
		return
	}
	tx = tx.Begin()
	taskinfo := &dao.TaskInfo{
		Id:          params.ID,
		Host:        params.Host,
		User:        params.User,
		DBName:      params.DBName,
		Password:    params.Password,
		BackupCycle: params.BackupCycle,
		KeepNumber:  params.KeepNumber,
		IsAllDBBak:  params.IsAllDBBak,
	}
	if err := taskinfo.Updates(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 30002, err)
		return
	}
	ding := &dao.DingDatabase{
		TaskID:          taskinfo.Id,
		DingSecret:      params.DingSecret,
		DingAccessToken: params.DingAccessToken,
		IsDingSend:      params.IsDingSend}
	if err := ding.Updates(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 30002, err)
		return
	}

	oss := &dao.OssDatabase{
		TaskID:     taskinfo.Id,
		IsOssSave:  params.IsOssSave,
		OssType:    params.OssType,
		Endpoint:   params.Endpoint,
		OssAccess:  params.OssAccess,
		OssSecret:  params.OssSecret,
		BucketName: params.BucketName,
		Directory:  params.Directory,
	}
	if err := oss.Updates(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 30002, err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(c, "任务更改成功")
}

func TaskPingCheck(task *dto.TaskAddInput) error {
	en, err := xorm.NewEngine("mysql", task.User+":"+task.Password+"@tcp("+task.Host+")/"+task.DBName+"?charset=utf8&parseTime=true")
	if err != nil {
		log.Logger.Errorf("创建数据库连接失败:%s", err)
		return err
	}
	if err = en.Ping(); err != nil {
		return err
	}
	return nil
}
