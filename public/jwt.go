package public

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"time"
)

var JWTToken jwtToken

// 定义jwtToken结构体
type jwtToken struct{}

// 从配置中获取jwtSecret
var jwtSecret = []byte("gin-mysqlbak")

// CustomClaims 自定义token中携带的信息
type CustomClaims struct {
	Uid int
	jwt.StandardClaims
}

// GenerateToken 生成token函数方法
func (*jwtToken) GenerateToken(uid *int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := CustomClaims{
		*uid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	log.Logger.Info("生成token信息成功!")
	return token, err
}

// ParseToken 解析token函数
func (*jwtToken) ParseToken(tokenString string) (claims *CustomClaims, err error) {
	// 使用jwt.ParseWithClaims方法解析token，这个token是前端传给我们的,获得一个*Token类型的对象
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("gin-mysqlbak"), nil
	})
	if err != nil {
		log.Logger.Error("解析token失败,错误信息," + err.Error())
		// 处理token解析后的各种错误
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token格式错误," + err.Error())
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token已过期," + err.Error())
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token还不可用," + err.Error())
			} else {
				return nil, errors.New("token不可用," + err.Error())
			}
		}
	}
	// 转换为*CustomClaims类型并返回
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		// 如果解析成功并且token是可用的
		return claims, nil
	}
	return nil, errors.New("解析token失败")
}
