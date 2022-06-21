package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbakv2/public"
	"time"
)

type AdminLoginInput struct {
	UserName string `form:"username" json:"username" comment:"用户名"  validate:"required,is_valid_username" example:"admin"`
	Password string `form:"password" json:"password" comment:"密码"   validate:"required" example:"123456"`
}

type AdminLoginOut struct {
	Token string `form:"token" json:"token" comment:"token"  example:"token"`
}

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	UserName  string    `json:"username"`
	LoginTime time.Time `json:"login_time"`
}

// BindValidParm 绑定并校验参数
func (a *AdminLoginInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}
