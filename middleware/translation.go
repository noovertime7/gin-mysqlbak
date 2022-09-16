package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/noovertime7/gin-mysqlbak/public"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"strings"
)

//设置Translation
func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//参照：https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go

		//设置支持语言
		en := en.New()
		zh := zh.New()

		//设置国际化翻译器
		uni := ut.New(zh, zh, en)
		val := validator.New()

		//根据参数取翻译器实例
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := uni.GetTranslator(locale)

		//翻译器注册到validator
		switch locale {
		case "en":
			err := en_translations.RegisterDefaultTranslations(val, trans)
			if err != nil {
				return
			}
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("en_comment")
			})
			break
		default:
			err := zh_translations.RegisterDefaultTranslations(val, trans)
			if err != nil {
				return
			}
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("comment")
			})

			//自定义验证方法
			//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
			//err = val.RegisterValidation("is_valid_username", func(fl validator.FieldLevel) bool {
			//	return fl.Field().String() == "admin"
			//})
			//if err != nil {
			//	return
			//}

			//自定义验证器
			//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
			//err = val.RegisterTranslation("is_valid_username", trans, func(ut ut.Translator) error {
			//	return ut.Add("is_valid_username", "{0} 填写不正确哦", true)
			//}, func(ut ut.Translator, fe validator.FieldError) string {
			//	t, _ := ut.T("is_valid_username", fe.Field())
			//	return t
			//})
			//if err != nil {
			//	return
			//}

			//ParseValidation
			//自定义验证方法
			_ = val.RegisterValidation("is_valid_bycle", func(fl validator.FieldLevel) bool {
				_, err = public.Cronexpr(fl.Field().String())
				if err != nil {
					return false
				}
				return true
			})
			//自定义验证器

			_ = val.RegisterTranslation("is_valid_bycle", trans, func(ut ut.Translator) error {
				return ut.Add("is_valid_bycle", "{0}格式不正确,正确格式: 30 12 * * ?", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("is_valid_bycle", fe.Field())
				return t
			})
			//host_valid
			//自定义验证方法
			_ = val.RegisterValidation("host_valid", func(fl validator.FieldLevel) bool {
				lenth := len(strings.Split(fl.Field().String(), ":"))
				if lenth != 2 {
					return false
				}
				return true
			})
			//自定义验证器

			_ = val.RegisterTranslation("host_valid", trans, func(ut ut.Translator) error {
				return ut.Add("host_valid", "{0}请添加端口信息,正确格式: 127.0.0.0:3306", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("host_valid", fe.Field())
				return t
			})
			break
		}
		c.Set(public.TranslatorKey, trans)
		c.Set(public.ValidatorKey, val)
		c.Next()
	}
}
