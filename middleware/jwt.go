package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/public"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
	"github.com/pkg/errors"
)

// JWTAuth jwt认证函数
func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		if len(context.Request.URL.String()) >= 18 && context.Request.Method == "POST" && context.Request.URL.String()[0:18] == "/admin_login/login" {
			context.Next()
			return
		}
		// 处理验证逻辑
		token := context.Request.Header.Get("Access-Token")
		if token == "" {
			ResponseError(context, globalError.NewGlobalError(globalError.AuthorizationError, errors.New("UnAuthorization")))
			context.Abort()
			return
		}
		// 解析token内容
		claims, err := public.JWTToken.ParseToken(token)
		if err != nil {
			// token过期错误
			if err.Error() == "TokenExpired" {
				ResponseError(context, globalError.NewGlobalError(globalError.AuthorizationError, err))
				context.Abort()
				return
			}
			// 解析其他错误
			ResponseError(context, globalError.NewGlobalError(globalError.AuthorizationError, err))
			context.Abort()
			return
		}
		context.Set("uid", claims.Uid)
		context.Set("token", token)
		context.Next()
	}
}
