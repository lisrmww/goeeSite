package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"goser/utils"
	"net/http"
)

type BaseCtrl struct {
}

func (b BaseCtrl) error(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		m := errs.Translate(utils.T)
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"msg":     getErrMsg(m),
			},
		)
	} else {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"msg":     err.Error(),
			},
		)
	}
}

func (b BaseCtrl) success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     msg,
		"data":    data,
	})
}

func getErrMsg(m map[string]string) string {
	var msg string
	var i int
	for _, v := range m {
		i++
		if i < len(m) {
			msg += fmt.Sprintf("%v&", v)
		} else {
			msg += fmt.Sprintf("%v", v)
		}
	}
	return msg
}
