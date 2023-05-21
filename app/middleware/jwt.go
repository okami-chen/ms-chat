package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/okamin-chen/service/pkg/global"
	"net/http"
	"strings"
)

func JwtMiddleware(c *gin.Context) {

	l := c.Copy()

	authHeader := l.Request.Header.Get("authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 2003,
			"msg":  "请求头中auth为空",
		})
		c.Abort()
		return
	}

	// 按空格分割
	parts := strings.Split(authHeader, ".")
	if len(parts) != 3 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 2004,
			"msg":  "请求头中auth格式有误",
		})
		c.Abort()
		return
	}

	token, _ := global.Conf.Jwt.Parse(authHeader)
	//
	//token, _ := jwt.ParseWithClaims(authHeader, config.Claims{}, func(token *jwt.Token) (interface{}, error) {
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("Unexpected Signing Method: %v", token.Header["alg"])
	//	}
	//	return []byte(global.Conf.Jwt.Key), nil
	//})

	if token != nil {
		global.Log.Infoln(token.Claims)
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 2005,
			"msg":  "无效的Token",
		})
		c.Abort()
		return
	}

}
