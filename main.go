package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"goser/dao/mysql"
	"goser/router"
	"goser/utils"
	"time"
)

/**
1. 加载配置
2. 初始化日志
3. 初始化MySQL连接
4. 初始化Redis连接
5. 注册路由
6. 启动服务（关机）
*/

func main() {

	//创建一个默认的路由引擎
	r := gin.Default()
	//配置gin允许跨域请求

	if err := utils.ValidatorTranslate("zh"); err != nil {
		return
	}

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		ExposeHeaders:    []string{"Content-Length"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	router.UsersRouter(r)
	defer mysql.CloseMysqlConn()
	if err := r.Run(":3009"); err != nil {
		fmt.Println("启动失败")
	}

}
