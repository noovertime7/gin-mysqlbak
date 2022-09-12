package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/dao"
	"github.com/noovertime7/gin-mysqlbak/dto"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/wumansgy/goEncrypt/rsa"
	"log"
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
	// 解密前端传递来的密码
	privateKey := `-----BEGIN RSA PRIVATE KEY-----
	MIICXQIBAAKBgQDAJRUhzFZ9P64cic8slOpn82VlYJUusLKWTKqugn7lgNVUpWdV
	CagfhtkViTUg5KRvpGrESmrQPlRiImm/iX/rOKtQSyMq6UroBpafjL3t6sPyHHFo
	hdDVakR7b6S6UG6ZOaTMOC/avPWtInXgzU05sHjRqEIGFapVejRtgfPtbwIDAQAB
	AoGBAJ+grxOrHNdlBhLzckhJVwwRK1WzjXyCk3tGKi5cf2vPQmvWFiiRozi94K+B
	k7/F885EO+bjJCXpAlWc3Vmgs8GYPIqUPJSTVkiDMf8nV5qicxzOH/pqye6KE1DD
	C/m3gPJTwOk/oT1KWsA4AHn8wmup01mDkMX82U9WFtJ1ZXXBAkEA5vTvXsFhaizK
	sAm8rkKjavxulCtuuNgyN+w68AUcdX92Cb4Hw6cWnAlUyqYncWLu8+3/TVWrtJWT
	vP0zcm1k8QJBANT6xzIRdLZQYPHODZDD0p575rfJgR4wuZi0tzhPmJKdndjfaxaW
	fco9GjL2yZzH4aNpF+ReN21RmT1ewgKy+F8CQDrEcHRH+KWvqBOLJrugsTxz5x9E
	vfPC72RTc9vHMSqkuEBaXldmmNYzeaPnC3pKlkrzcFcZSYu109XvB7xCIcECQDiE
	P73SkgUbOU6RXlovDMIPoP7eUwwe4/FY61HfFV66wrtdNj6tOr4jDsO9Z2zaQc8q
	QTPRqKWyxJZbgeJTecMCQQDFnb5C0dHqpDE1PBHklzo9TnycUG9T1gBIVC7oZDex
	ImeHIUC/olM27UPRzf4ku+ZtMb+bTZpjcUcRBzs5JnAK
	-----END RSA PRIVATE KEY-----`
	ciphertext, err := base64.StdEncoding.DecodeString(params.Password)
	if err != nil {
		log.Println(err)
		return
	}
	plaintext, err := rsa.RsaDecryptByHex(string(ciphertext), privateKey)
	if err != nil {
		log.Println("解密失败", err)
		return
	}
	fmt.Println(string(plaintext))
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
