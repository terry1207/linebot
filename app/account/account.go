package account

import (
	"fmt"
	"linebot/internal/errmsg"
	"linebot/internal/model/account"

	"linebot/internal/response"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var account account.Account
	err := c.BindJSON(&account)
	if err != nil {
		return
	}

	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()

	err = Add(&account)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	fmt.Printf("%v\n", account)
	c.JSON(200, account)
	response.Response(c, errmsg.SUCCESS)
}

func GetAccount(c *gin.Context) {

	accounts, err := GetAll()

	if err != nil {
		response.Response(c, errmsg.ERROR_ACCOUNT_NOT_EXIST)
		return
	}

	c.JSON(http.StatusOK, accounts)
}

func GetAccountById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("AccountId"))
	account, err := GetById(int64(id))
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	c.JSON(200, account)
}

func UpdateAccount(c *gin.Context) {

}

func DeleteAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("accountId"))

	account, err := DeleteById(int64(id))

	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	c.JSON(200, account)
}
