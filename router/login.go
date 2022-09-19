package router

import (
	"github.com/gin-gonic/gin"
	"goser/controllers"
)

func UsersRouter(engine *gin.Engine) {
	r := engine.Group("/user")
	{
		// captcha
		r.GET("/captcha", controllers.UserCtrl{}.Captcha)
		// login
		r.POST("/login", controllers.UserCtrl{}.Login)
	}
}
