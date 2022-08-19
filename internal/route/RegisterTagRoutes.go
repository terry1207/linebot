package route

import (
	v1 "linebot/api/v1"
	"linebot/internal/middleware"

	"github.com/gin-gonic/gin"
)

var RegisterTagRoutes = func(r *gin.Engine) {
	r.POST("/tag", middleware.JwtMiddleware(), v1.CreateTag)
	r.GET("/tag", middleware.JwtMiddleware(), v1.GetTag)
	r.GET("/tag/:tagId", middleware.JwtMiddleware(), v1.GetTagById)
	r.PUT("/tag/:tagId", middleware.JwtMiddleware(), v1.UpdateTag)
	r.DELETE("/tag/:tagId", middleware.JwtMiddleware(), v1.DeleteTag)

}
