package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type TaskInfo struct {
	Id          int       `gorm:"primary_key" description:"自增主键"`
	HostID      int       `json:"id" gorm:"column:host_id" description:"主机关系id"`
	DBName      string    `json:"db_name" gorm:"column:db_name" description:"备份库名"`
	BackupCycle string    `json:"backup_cycle" gorm:"column:backup_cycle" description:"备份周期"`
	KeepNumber  int       `json:"keep_number"  gorm:"column:keep_number" description:"数据保留周期"`
	IsAllDBBak  int       `json:"is_all_dbbak" gorm:"column:is_all_dbbak" description:"是否全库备份"`
	IsDelete    int       `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
	Status      int       `json:"status" gorm:"column:status" description:"开关"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at" description:"更新时间"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at" description:"添加时间"`
}

func (s *TaskInfo) TableName() string {
	return "t_taskinfo"
}

func (s *TaskInfo) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(s).Error
}

func (s *TaskInfo) Find(c *gin.Context, tx *gorm.DB, search *TaskInfo) (*TaskInfo, error) {
	out := &TaskInfo{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *TaskInfo) FindAllTask(c *gin.Context, tx *gorm.DB, params *dto.HostIDInput) ([]*TaskInfo, error) {
	var result []*TaskInfo
	if params == nil {
		err := tx.WithContext(c).Where("is_delete = 0").Find(&result).Error
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	err := tx.WithContext(c).Where("is_delete = 0 and host_id = ?", params.HostID).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindAllStatusUpTask(tx *gorm.DB) ([]*TaskInfo, error) {
	var result []*TaskInfo
	err := tx.Where("is_delete = 0 and status = 1 ").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *TaskInfo) Updates(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Where("id = ?", d.Id).Updates(d).Error
}

// UpdatesStatus 只更新单个字段
func (d *TaskInfo) UpdatesStatus(tx *gorm.DB) error {
	if d.Id == 0 {
		return errors.New("TASK_INFO的ID为空,更新失败")
	}
	return tx.Table(d.TableName()).Where("id = ?", d.Id).Updates(map[string]interface{}{
		"status": d.Status,
	}).Error
}

func (s *TaskInfo) PageList(c *gin.Context, tx *gorm.DB, params *dto.TaskListInput) ([]TaskInfo, int, error) {
	var total int64 = 0
	list := []TaskInfo{}
	offset := (params.PageNo - 1) * params.PageSize
	query := tx.WithContext(c)
	query.Find(&list).Count(&total)
	if params.HostId > 0 {
		query = query.Table(s.TableName()).Where("is_delete=0 and host_id = ?", params.HostId)
	} else {
		query = query.Table(s.TableName()).Where("is_delete=0")
	}
	if params.Info != "" {
		query = query.Where("( db_name like ?)", "%"+params.Info+"%")
	}
	if err := query.Limit(params.PageSize).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return list, int(total), nil
}

func (s *TaskInfo) TaskDetail(c *gin.Context, tx *gorm.DB, serch *TaskInfo) (*TaskDetail, error) {
	info := &TaskInfo{Id: serch.Id}
	infores, err := info.Find(c, tx, info)
	if err != nil {
		return nil, err
	}
	hostinfo := &HostDatabase{Id: infores.HostID}
	hostinfores, err := hostinfo.Find(c, tx, hostinfo)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	ding := &DingDatabase{TaskID: serch.Id}
	dingres, err := ding.Find(c, tx, ding)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	oss := &OssDatabase{TaskID: serch.Id}
	ossres, err := oss.Find(c, tx, oss)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &TaskDetail{
		Host: hostinfores,
		Info: infores,
		Oss:  ossres,
		Ding: dingres,
	}, nil
}

func (t *TaskInfo) GroupByHost(c *gin.Context, tx *gorm.DB) ([]dto.DashServiceStatItemOutput, error) {
	var list []dto.DashServiceStatItemOutput
	query := tx.Table(t.TableName())
	if err := query.Where("is_delete=0").Select("host_id, count(*) as value").Group("host_id").Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
