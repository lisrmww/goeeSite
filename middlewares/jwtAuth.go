package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goser/utils"
	"net/http"
	"strings"
)

//路由白名单
var routerWhiteList = []string{"/user/login", "/user/logout"}

var code = 3000

func JwtAuth(c *gin.Context) {
	uri := c.Request.RequestURI
	tokenString := c.Request.Header.Get("Authorization")
	for _, r := range routerWhiteList {
		if r == uri {
			c.Next()
			return
		}
	}
	if len(tokenString) > 0 {
		authorization := strings.Split(tokenString, " ")
		token, claims, err := utils.ParseToken(authorization[1])
		if err != nil || !token.Valid {
			errType, _ := err.(*jwt.ValidationError)
			if errType.Errors == jwt.ValidationErrorExpired {
				code = 3001
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    code,
				"message": "token验证失败",
				"success": false,
			})
			c.Abort()
			return
		}
		if claims.ID != 0 {
			c.Set("id", claims.ID)
		}
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "无访问权限",
			"success": false,
		})
		c.Abort()
	}
}
