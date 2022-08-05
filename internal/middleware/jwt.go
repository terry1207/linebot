package middleware

import (
	"fmt"
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
		fmt.Println("authHeader", authHeader)
		c.Set("email", "eeeee")
		c.Next()
	}
}
