package route

import (
	"fmt"
	v1 "linebot/api/v1"
	"net/http"

	"linebot/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//richmenu.Build_RichMenu()

	router := gin.Default()
	router.Use(gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))
	router.LoadHTMLGlob("web/templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	router.GET("/repeat", v1.RepeatHandler)
	router.GET("/db/create", v1.CreateCampInfo)
	router.GET("db/get", v1.GetCampInfo)
	router.Any("/callback", v1.ReplyMessage)

	router.GET("/ce", middleware.MiddlewareTest(), func(c *gin.Context) {
		req, _ := c.Get("request")
		fmt.Println("request", req)
		c.JSON(200, gin.H{"request": req})
	})
	// lineroute := router.Group("/callback")
	// lineroute.Any("/", v1.ReplyMessage)

	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.

	return router
}
