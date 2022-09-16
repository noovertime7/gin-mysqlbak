package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/services"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type AdminController struct {
}

func AdminRegister(group *gin.RouterGroup) {
	admininfo := &AdminController{}
	group.GET("/role_info", admininfo.RoleInfo)
	group.GET("/user_info", admininfo.UserInfo)
	group.POST("/changepwd", admininfo.ChangePwd)
}

func (a *AdminController) UserInfo(ctx *gin.Context) {
	//1、通过claims解析用户id
	claims, exists := ctx.Get("claims")
	if !exists {
		log.Logger.Error("claims不存在,请检查jwt中间件")
	}
	cla, _ := claims.(*public.CustomClaims)
	userinfo, err := services.UserService.GetUserInfo(ctx, cla.Uid)
	if err != nil {
		log.Logger.Error("查询用户信息失败", err)
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	middleware.ResponseSuccess(ctx, userinfo)
}

func (a *AdminController) RoleInfo(ctx *gin.Context) {
	//1、通过claims解析用户id
	claims, exists := ctx.Get("claims")
	if !exists {
		log.Logger.Error("claims不存在,请检查jwt中间件")
	}
	cla, _ := claims.(*public.CustomClaims)
	//2、取出数据然后封装输出
	roleInfo, err := services.RuleService.GetRoleInfo(ctx, cla.Uid)
	if err != nil {
		log.Logger.Error("查询用户权限失败", err)
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	userinfo, err := services.UserService.GetUserInfo(ctx, cla.Uid)
	if err != nil {
		log.Logger.Error("查询用户信息失败", err)
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	out := &dto.AdminInfoOutput{
		UserInfoOutPut: userinfo,
		Role:           roleInfo,
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
	if err := services.UserService.ChangePwd(ctx, params); err != nil {
		log.Logger.Error("修改密码失败", err)
		middleware.ResponseError(ctx, 30002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "更改密码成功")
}
