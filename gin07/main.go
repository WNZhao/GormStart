package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translator "github.com/go-playground/validator/v10/translations/en"
	zh_translator "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)

// 绑定为 json
type LoginForm struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type RegisterForm struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 跨字段了
}

var trans ut.Translator

// 翻译
func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT)
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("locale %s not found", locale)
		}
		switch locale {
		case "en":
			en_translator.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translator.RegisterDefaultTranslations(v, trans)
		default:
			en_translator.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func main() {
	if err := InitTrans("zh"); err != nil {
		fmt.Println(err.Error())
		return
	}
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if err := c.ShouldBind(&form); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": errs.Translate(trans)})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
	})
	r.POST("/register", func(c *gin.Context) {
		var signUpForm RegisterForm
		if err := c.ShouldBind(&signUpForm); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": errs.Translate(trans)})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})
	r.Run(":8083")
}
