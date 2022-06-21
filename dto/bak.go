package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
)

type Bak struct {
	ID int `json:"id" form:"id" validate:"required,min=1"`
}

func (d *Bak) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, d)
}
