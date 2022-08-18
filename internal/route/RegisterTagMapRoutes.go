package route

import (
	v1 "linebot/api/v1"
	"linebot/internal/middleware"

	"github.com/gin-gonic/gin"
)

var RegisterTagMapRoutes = func(r *gin.Engine) {
	r.POST("/camp", middleware.JwtMiddleware(), v1.CreateTagMap)
	r.GET("/camp", middleware.JwtMiddleware(), v1.GetTagMap)
	r.GET("/camp/:campId", middleware.JwtMiddleware(), v1.GetTagMapById)
	r.PUT("/camp/:campId", middleware.JwtMiddleware(), v1.UpdateTagMap)
	r.DELETE("/camp/:campId", middleware.JwtMiddleware(), v1.DeleteTagMap)

}
