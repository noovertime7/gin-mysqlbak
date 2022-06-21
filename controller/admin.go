package controller

import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
)

type AdminController struct {
}

func AdminRegister(group *gin.RouterGroup) {
	admininfo := &AdminController{}
	group.GET("/admin_info", admininfo.AdminInfo)
	group.POST("/changepwd", admininfo.ChangePwd)

}

func (a *AdminController) AdminInfo(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	sessinfo := sess.Get(public.AdminSessionInfoKey)
	adminsessioninfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessinfo)), adminsessioninfo); err != nil {
		middleware.ResponseError(ctx, 3000, err)
		return
	}
	//1、读取sessionkey对应的json，转换为结构体
	//2、取出数据然后封装输出
	out := &dto.AdminInfoOutput{
		ID:           adminsessioninfo.ID,
		Name:         adminsessioninfo.UserName,
		LoginTime:    adminsessioninfo.LoginTime,
		Avatar:       "https://img2.woyaogexing.com/2022/06/19/b0be13535a86d0b0!400x400.jpg",
		Introduction: "我是admin",
		Roles:        []string{"admin"},
	}
	middleware.ResponseSuccess(ctx, out)
}

// 更改密码
func (a *AdminController) ChangePwd(ctx *gin.Context) {
	params := &dto.ChangePwdInput{}
	if err := params.BindValidParm(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}
	//1、从session读取用户信息到结构体 sessionInfo
	sess := sessions.Default(ctx)
	sessinfo := sess.Get(public.AdminSessionInfoKey)
	adminsessioninfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessinfo)), adminsessioninfo); err != nil {
		middleware.ResponseError(ctx, 3000, err)
		return
	}
	//2、利用结构体中的id去读取数据库信息 adminInfo
	//获取数据库连接池
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	adminInfo := &dao.Admin{}
	adminInfo, err = adminInfo.Find(ctx, tx, &dao.Admin{UserName: adminsessioninfo.UserName})
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
