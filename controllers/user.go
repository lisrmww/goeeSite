package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"goser/dao/mysql"
	"goser/models"
	"goser/utils"
	"net/http"
)

type UserCtrl struct {
	BaseCtrl
	AccessToken string `json:"accessToken"`
}

type LoginForm struct {
	Username     string `json:"username" form:"username" binding:"required,min=5,max=10"`
	Password     string `json:"password" form:"password" binding:"required,min=6,max=16"`
	CaptchaId    string `json:"captchaId" form:"captchaId" binding:"required"`
	CaptchaValue string `json:"captchaValue" form:"captchaValue" binding:"required"`
}

// Captcha
// 获取验证码
func (u UserCtrl) Captcha(c *gin.Context) {
	captchaId, captchaImage, err := utils.MakeCaptcha()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    captchaId,
		"captchaImage": captchaImage,
	})
}

// Login
// 登录
func (u UserCtrl) Login(c *gin.Context) {
	var loginForm LoginForm
	if err := c.ShouldBind(&loginForm); err != nil {
		u.error(c, err)
		return
	}
	vc := utils.VerifyCaptcha(loginForm.CaptchaId, loginForm.CaptchaValue)
	if !vc {
		u.error(c, errors.New("验证码错误"))
		return
	}
	var user []models.User
	mysql.MySQL.Where("username=? AND password=?", loginForm.Username, loginForm.Password).First(&user)
	if len(user) > 0 {
		accessToken, refreshToken, err := utils.GenerateToken(user[0].ID)
		var token = map[string]string{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		}
		if err != nil {
			fmt.Println(err.Error())
		}
		u.success(c, "登录成功", token)
	} else {
		u.error(c, errors.New("用户名或密码错误"))
	}
}
