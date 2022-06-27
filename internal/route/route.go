package route

import (
	v1 "linebot/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//front.Init()

	router := gin.New()
	// router.Use(gin.Logger())
	// router.LoadHTMLGlob("web/templates/*.tmpl.html")
	// router.Static("/static", "static")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })
	//router.GET("/repeat", v1.RepeatHandler)
	router.GET("/db", v1.DbTest)
	router.Any("/callback", v1.ReplyMessage)

	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.

	return router
}
