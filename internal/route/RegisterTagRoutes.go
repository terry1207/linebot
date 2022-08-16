package route

import (
	v1 "linebot/api/v1"
	"linebot/internal/middleware"

	"github.com/gin-gonic/gin"
)

var RegisterTagRoutes = func(r *gin.Engine) {
	r.POST("/camp", middleware.JwtMiddleware(), v1.CreateTag)
	r.GET("/camp", middleware.JwtMiddleware(), v1.GetTag)
	r.GET("/camp/:campId", middleware.JwtMiddleware(), v1.GetTagById)
	r.PUT("/camp/:campId", middleware.JwtMiddleware(), v1.UpdateTag)
	r.DELETE("/camp/:campId", middleware.JwtMiddleware(), v1.DeleteTag)

}
