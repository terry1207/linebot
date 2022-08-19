package route

import (
	v1 "linebot/api/v1"
	"linebot/internal/middleware"

	"github.com/gin-gonic/gin"
)

var RegisterTagMapRoutes = func(r *gin.Engine) {
	r.POST("/tagmap", middleware.JwtMiddleware(), v1.CreateTagMap)
	r.GET("/tagmap", middleware.JwtMiddleware(), v1.GetTagMap)
	r.GET("/tagmap/:tagmapId", middleware.JwtMiddleware(), v1.GetTagMapById)
	r.PUT("/tagmap/:tagmapId", middleware.JwtMiddleware(), v1.UpdateTagMap)
	r.DELETE("/tagmap/:tagmapId", middleware.JwtMiddleware(), v1.DeleteTagMap)

}
