package utils

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store = base64Captcha.DefaultMemStore

// MakeCaptcha 获取验证码
func MakeCaptcha() (id string, b64s string, err error) {
	var driver base64Captcha.Driver
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      1,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "0123456789qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 65,
			G: 102,
			B: 214,
			A: 90,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driver = driverString.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = c.Generate()
	return id, b64s, err
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(id string, VerifyValue string) bool {
	if store.Verify(id, VerifyValue, true) {
		return true
	} else {
		return false
	}
}
