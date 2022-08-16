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

func CreateCamp(c *gin.Context) {
	var camp repository.Camp
	err := c.BindJSON(&camp)
	if err != nil {
		return
	}

	camp.CreatedAt = time.Now()
	camp.UpdatedAt = time.Now()
	fmt.Printf("%v\n", camp)

	err = camp.CreateNewCamp()
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	fmt.Printf("%v\n", camp)
	c.JSON(200, camp)
	response.Response(c, errmsg.SUCCESS)
}

func GetCamp(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("campId"))

	camp, err := repository.GetCampById(int64(id))
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	if camp.CampName == "" {
		response.Response(c, errmsg.ERROR_ACCOUNT_NOT_EXIST)
		return
	}

	c.JSON(http.StatusOK, camp)
}

func GetCampById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("campId"))
	c.String(200, "%s", id)
}

func UpdateCamp(c *gin.Context) {

}

func DeleteCamp(c *gin.Context) {

}
