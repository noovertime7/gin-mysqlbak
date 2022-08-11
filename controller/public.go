package controller

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/conf"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/pkg/errors"
	"strings"
)

type PublicController struct{}

func PublicRegister(group *gin.RouterGroup) {
	pb := &PublicController{}
	group.GET("/download", pb.DownLoadBakfile)
	group.GET("/check_file_exists", pb.BakFileExists)
}

func (p *PublicController) DownLoadBakfile(ctx *gin.Context) {
	params := &dto.Bak{}
	if err := params.BindValidParm(ctx); err != nil {
		log.Logger.Error(err)
		return
	}
	tx, _ := lib.GetGormPool("default")
	bakhistory := &dao.BakHistory{
		Id: params.ID,
	}
	resBakHistory, err := bakhistory.Find(ctx, tx, bakhistory)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	filepath := resBakHistory.FileName
	filename := strings.Split(filepath, "/")[len(strings.Split(filepath, "/"))-1]
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	////ctx.Header("Content-Disposition", "inline;filename="+filename)
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Cache-Control", "no-cache")
	ctx.File(filepath)
}

func (p *PublicController) BakFileExists(ctx *gin.Context) {
	params := &dto.Bak{}
	if err := params.BindValidParm(ctx); err != nil {
		log.Logger.Error(err)
		return
	}
	tx, _ := lib.GetGormPool("default")
	bakhistory := &dao.BakHistory{
		Id: params.ID,
	}
	resBakHistory, err := bakhistory.Find(ctx, tx, bakhistory)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	filepath := resBakHistory.FileName
	if ok, _ := public.HasDir(filepath); !ok {
		middleware.ResponseError(ctx, 20001, errors.New("本地文件不存在"))
		return
	}
	clusterUrl := conf.GetStringConf("base", "download_url")
	middleware.ResponseSuccess(ctx, clusterUrl+"/public/download")
}
