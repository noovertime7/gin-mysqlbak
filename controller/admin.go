package controller

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/services"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"time"
)

type AdminController struct {
}

func AdminRegister(group *gin.RouterGroup) {
	admininfo := &AdminController{}
	group.GET("/admin_info", admininfo.AdminInfo)
	group.POST("/changepwd", admininfo.ChangePwd)
}

func (a *AdminController) AdminInfo(ctx *gin.Context) {
	//1、通过claims解析用户id
	claims, exists := ctx.Get("claims")
	if !exists {
		log.Logger.Error("claims不存在,请检查jwt中间件")
	}
	cla, _ := claims.(*public.CustomClaims)
	//从数据库查询用户信息
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	adminDB := &dao.Admin{Id: cla.Uid}
	admin, err := adminDB.Find(ctx, tx, adminDB)
	//2、取出数据然后封装输出
	roleInfo, err := services.RuleService.GetRoleInfo(ctx, cla.Uid)
	if err != nil {
		return
	}
	out := dto.AdminInfoOutput{
		ID:           admin.Id,
		Name:         admin.UserName,
		LoginTime:    time.Now(),
		Avatar:       "",
		Introduction: "用户介绍",
		Status:       admin.Status,
		CreatorId:    "system",
		Role:         roleInfo,
	}
	middleware.ResponseSuccess(ctx, out)
}

// ChangePwd 更改密码
func (a *AdminController) ChangePwd(ctx *gin.Context) {
	params := &dto.ChangePwdInput{}
	if err := params.BindValidParm(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}
	//1、通过claims解析用户id
	claims, exists := ctx.Get("claims")
	if !exists {
		log.Logger.Error("claims不存在,请检查jwt中间件")
	}
	cla, _ := claims.(*public.CustomClaims)
	//2、利用结构体中的id去读取数据库信息 adminInfo
	//获取数据库连接池
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	adminInfo := &dao.Admin{}
	adminInfo, err = adminInfo.Find(ctx, tx, &dao.Admin{Id: cla.Uid})
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	//3、加盐 params.Password + admininfo.salt sha256 saltPassword
	saltPassword := public.GenSaltPassword(adminInfo.Salt, params.Password)
	//4、保存新的password到数据库中
	adminInfo.Password = saltPassword
	if err := adminInfo.Save(ctx, tx); err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "更改密码成功")
}
