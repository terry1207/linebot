package route

import (
	"linebot/app/account"
	"linebot/internal/middleware"

	"github.com/gin-gonic/gin"
)

var RegisterAccountRoutes = func(r *gin.Engine) {
	r.POST("/user", middleware.JwtMiddleware(), account.CreateAccount)
	r.GET("/user", middleware.JwtMiddleware(), account.GetAccount)
	r.GET("/user/:userId", middleware.JwtMiddleware(), account.GetAccountById)
	r.PUT("/user/:userId", middleware.JwtMiddleware(), account.UpdateAccount)
	r.DELETE("/user/:userId", middleware.JwtMiddleware(), account.DeleteAccount)

}
