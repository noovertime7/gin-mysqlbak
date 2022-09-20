package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/services"
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
	if err := params.BindValidParams(ctx); err != nil {
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
	//log.Logger.Debug("解密后密码:", params.Password)
	token, err := services.UserService.Login(ctx, params)
	if err != nil {
		log.Logger.Error("登录失败", err)
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	out := &dto.AdminLoginOut{Token: token, Message: "登录成功"}
	middleware.ResponseSuccess(ctx, out)
}

func (a *AdminLoginController) AdminLoginOut(ctx *gin.Context) {
	if err := services.UserService.LoginOut(ctx); err != nil {
		log.Logger.Error("推出登录失败:", err)
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "退出成功")
}
