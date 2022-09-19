package utils

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var T ut.Translator

//ValidatorTranslate
//验证信息翻译器
func ValidatorTranslate(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		//获取标签名
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			return FirstLower(field.Name)
		})

		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT, enT)

		var ok bool
		T, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, T)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, T)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, T)
		}

		return nil
	}
	return nil
}
