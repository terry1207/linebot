package route

import (
	v1 "linebot/api/v1"
	"linebot/internal/middleware"

	"github.com/gin-gonic/gin"
)

var RegisterCampRoutes = func(r *gin.Engine) {
	r.POST("/camp", v1.CreateCamp)
	r.GET("/camp", middleware.JwtMiddleware(), v1.GetCamp)
	r.GET("/camp/:campId", middleware.JwtMiddleware(), v1.GetCampById)
	r.PUT("/camp/:campId", middleware.JwtMiddleware(), v1.UpdateCamp)
	r.DELETE("/camp/:campId", middleware.JwtMiddleware(), v1.DeleteCamp)

}
