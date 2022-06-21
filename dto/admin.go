package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
	"time"
)

type AdminInfoOutput struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}

type ChangePwdInput struct {
	Password string `form:"password" json:"password" comment:"密码"   validate:"required" example:"123456"`
}

func (a *ChangePwdInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}
