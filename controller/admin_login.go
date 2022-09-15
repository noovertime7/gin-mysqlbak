package controller

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type AdminLoginController struct {
}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
	group.GET("/loginout", adminLogin.AdminLoginOut)
}

// AdminLogin godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param polygon body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOut} "success"
// @Router /admin_login/login [post]
func (a *AdminLoginController) AdminLogin(ctx *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParm(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}
	var err error
	params.Password, err = public.RsaDecode(params.Password)
	if err != nil {
		log.Logger.Error("rsa解密失败", err)
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	log.Logger.Debug("解密后密码:", params.Password)
	//获取数据库连接池
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	//进行密码校验
	admin := &dao.Admin{}
	admin, err = admin.LoginCheck(ctx, tx, params)
	if err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	//生成token
	token, err := public.JWTToken.GenerateToken(&admin.Id)
	if err != nil {
		log.Logger.Error("生成token失败", err)
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	//更新用户在线状态
	admin.Status = 1
	if err = admin.UpdateStatus(ctx, tx, admin); err != nil {
		log.Logger.Error("更新用户状态失败", err)
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	out := &dto.AdminLoginOut{Token: token, Message: "登录成功"}
	middleware.ResponseSuccess(ctx, out)
}

func (a *AdminLoginController) AdminLoginOut(ctx *gin.Context) {
	//获取数据库连接池
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	claims, exists := ctx.Get("claims")
	if !exists {
		log.Logger.Error("claims不存在,请检查jwt中间件")
	}
	cla, _ := claims.(*public.CustomClaims)
	adminDB := &dao.Admin{Id: cla.Uid}
	admin, err := adminDB.Find(ctx, tx, adminDB)
	admin.Status = 0
	if err := admin.UpdateStatus(ctx, tx, admin); err != nil {
		log.Logger.Error("退出失败", err)
		middleware.ResponseError(ctx, 2000, err)
		return
	}
	middleware.ResponseSuccess(ctx, "退出成功")
}
