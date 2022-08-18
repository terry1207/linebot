package middleware

import (
	"fmt"
	"linebot/internal/errmsg"
	"linebot/internal/response"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddlewareTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		fmt.Println("執行middleware")

		c.Set("request", "middle")
		status := c.Writer.Status()

		t2 := time.Since(t)
		fmt.Println("time", t2, "middle ware finish", status)

		c.Next()

		fmt.Println("api finish", "time", time.Since(t))

	}
}

func JwtMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Response(c, errmsg.TOKEN_NOT_FOUND)

			//執行完當前的handler 其餘皆不會執行
			c.Abort()
			return
		}

		//按照空格分割
		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Response(c, errmsg.TOKEN_FORMAT_ERROR)
			c.Abort()
			return
		}

		c.Set("token", parts[1])

		c.Next()
	}
}
