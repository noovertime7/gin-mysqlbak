package controller

import (
	"encoding/json"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"time"
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
	//设置session
	sessioninin := &dto.AdminSessionInfo{
		ID:        admin.Id,
		UserName:  admin.UserName,
		LoginTime: time.Now(),
	}
	sessBts, err := json.Marshal(sessioninin)
	if err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	sess := sessions.Default(ctx)
	sess.Set(public.AdminSessionInfoKey, string(sessBts))
	if err := sess.Save(); err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	out := &dto.AdminLoginOut{Token: admin.UserName}
	middleware.ResponseSuccess(ctx, out)
}

func (a *AdminLoginController) AdminLoginOut(ctx *gin.Context) {
	//获取数据库连接池
	sess := sessions.Default(ctx)
	sess.Delete(public.AdminSessionInfoKey)
	if err := sess.Save(); err != nil {
		middleware.ResponseError(ctx, 3002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "退出成功")
}
