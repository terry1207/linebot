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
	var new_account account.Account
	err := c.BindJSON(&new_account)
	if err != nil {
		return
	}

	new_account.CreatedAt = time.Now()
	new_account.UpdatedAt = time.Now()

	err = account.Add(&new_account)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	fmt.Printf("%v\n", new_account)
	c.JSON(200, new_account)
	response.Response(c, errmsg.SUCCESS)
}

func GetAccount(c *gin.Context) {

	accounts, err := account.GetAll()

	if err != nil {
		response.Response(c, errmsg.ERROR_ACCOUNT_NOT_EXIST)
		return
	}

	c.JSON(http.StatusOK, accounts)
}

func GetAccountById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("AccountId"))
	account, err := account.GetById(int64(id))
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

	account, err := account.DeleteById(int64(id))

	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	c.JSON(200, account)
}
