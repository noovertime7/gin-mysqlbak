package middleware

import (
	"errors"
	"fmt"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
	"runtime/debug"
)

// RecoveryMiddleware RecoveryMiddleware捕获所有panic，并且返回错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//先做一下日志记录
				fmt.Println(string(debug.Stack()))
				public.ComLogNotice(c, "_com_panic", map[string]interface{}{
					"error": fmt.Sprint(err),
					"stack": string(debug.Stack()),
				})

				if lib.ConfBase.DebugMode != "debug" {
					ResponseError(c, globalError.NewGlobalError(globalError.ServerError, errors.New("")))
					return
				} else {
					ResponseError(c, globalError.NewGlobalError(globalError.AuthorizationError, errors.New("")))
					return
				}
			}
		}()
		c.Next()
	}
}
