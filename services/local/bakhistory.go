package local

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/public/database"
	"sync"
)

var (
	localBakHistoryService *bakHistoryService
	bakHistoryServiceOnce  sync.Once
)

// GetBakHistoryService 单例模式
func GetBakHistoryService() *bakHistoryService {
	bakHistoryServiceOnce.Do(func() {
		localBakHistoryService = &bakHistoryService{}
	})
	return localBakHistoryService
}

type bakHistoryService struct{}

func (b *bakHistoryService) Delete(ctx *gin.Context, hid int) error {
	tx := database.GetDB()
	h := &dao.BakHistory{Id: hid}
	history, err := h.Find(ctx, tx, h)
	if err != nil {
		return err
	}
	history.IsDelete = sql.NullInt32{Int32: 1, Valid: true}
	return history.Updates(ctx, tx)
}

// GetHistoryList 用于获取历史记录列表
func (b *bakHistoryService) GetHistoryList(ctx *gin.Context, params *dto.HistoryListInput) (*dto.HistoryListOutput, error) {
	his := &dao.BakHistory{}
	list, total, err := his.PageList(ctx, database.GetDB(), params)
	if err != nil {
		return nil, err
	}
	var outList []dto.HistoryListOutItem
	for _, listIterm := range list {
		outItem := dto.HistoryListOutItem{
			ID:         listIterm.Id,
			Host:       listIterm.Host,
			DBName:     listIterm.DBName,
			DingStatus: listIterm.DingStatus,
			OSSStatus:  listIterm.OssStatus,
			Message:    listIterm.Msg,
			FileSize:   listIterm.FileSize,
			FileName:   listIterm.FileName,
			BakTime:    listIterm.BakTime.Format("2006年01月02日15:04:01"),
		}
		outList = append(outList, outItem)
	}
	out := &dto.HistoryListOutput{
		Total:    total,
		List:     outList,
		PageSize: params.PageSize,
		PageNo:   params.PageNo,
	}
	return out, err
}

// GetHistoryNumInfo 用于统计备份文件个数，向前端展示
func (b *bakHistoryService) GetHistoryNumInfo(ctx *gin.Context) (*dto.HistoryNumInfoOutput, error) {
	tx := database.GetDB()
	his := &dao.BakHistory{}
	list, total, err := his.PageList(ctx, tx, &dto.HistoryListInput{
		Info:     "",
		PageNo:   1,
		PageSize: 99999,
	})
	//获取文件大小
	var allSize int
	for _, h := range list {
		allSize += h.FileSize
	}
	mbAllSize := fmt.Sprintf("%.2fMB", float64(allSize)/float64(1024))
	//获取一周内任务数
	dataList, err := his.FindByDate(ctx, tx, 7)
	if err != nil {
		return nil, err
	}
	return &dto.HistoryNumInfoOutput{
		WeekNums:    len(dataList),
		AllNums:     total,
		AllFileSize: mbAllSize,
	}, nil
}
