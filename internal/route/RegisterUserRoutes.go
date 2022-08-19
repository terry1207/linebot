package route

import (
	v1 "linebot/api/v1"
	"linebot/internal/middleware"

	"github.com/gin-gonic/gin"
)

var RegisterUserRoutes = func(r *gin.Engine) {
	r.POST("/user", middleware.JwtMiddleware(), v1.CreateUser)
	r.GET("/user", middleware.JwtMiddleware(), v1.GetUser)
	r.GET("/user/:userId", middleware.JwtMiddleware(), v1.GetUserById)
	r.PUT("/user/:userId", middleware.JwtMiddleware(), v1.UpdateUser)
	r.DELETE("/user/:userId", middleware.JwtMiddleware(), v1.DeleteUser)

}
