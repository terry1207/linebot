package route

import (
	v1 "linebot/api/v1"
	"linebot/internal/middleware"

	"github.com/gin-gonic/gin"
)

var RegisterUserRoutes = func(r *gin.Engine) {
	r.POST("/camp", middleware.JwtMiddleware(), v1.CreateUser)
	r.GET("/camp", middleware.JwtMiddleware(), v1.GetUser)
	r.GET("/camp/:campId", middleware.JwtMiddleware(), v1.GetUserById)
	r.PUT("/camp/:campId", middleware.JwtMiddleware(), v1.UpdateUser)
	r.DELETE("/camp/:campId", middleware.JwtMiddleware(), v1.DeleteUser)

}
