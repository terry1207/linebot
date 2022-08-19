package v1

import (
	"fmt"
	"linebot/internal/errmsg"
	"linebot/internal/repository"
	"linebot/internal/response"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user repository.User
	err := c.BindJSON(&user)
	if err != nil {
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = user.CreateNewUser()
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	fmt.Printf("%v\n", user)
	c.JSON(200, user)
	response.Response(c, errmsg.SUCCESS)
}

func GetUser(c *gin.Context) {

	users, err := repository.GetAllUser()

	if err != nil {
		response.Response(c, errmsg.ERROR_ACCOUNT_NOT_EXIST)
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("UserId"))
	user, err := repository.GetUserById(int64(id))
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("userId"))

	user, err := repository.DeleteUserById(int64(id))

	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	c.JSON(200, user)
}
