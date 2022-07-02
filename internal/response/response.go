package response

import (
	"net/http"

	"linebot/internal/errmsg"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, status int) {
	var context = gin.H{
		"status":  status,
		"message": errmsg.GetErrMsg(status),
	}
	c.JSON(
		http.StatusOK,
		context,
	)
}
