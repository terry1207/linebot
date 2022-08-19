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

func CreateCampRound(c *gin.Context) {
	var campRound repository.CampRound
	err := c.BindJSON(&campRound)
	if err != nil {
		return
	}

	campRound.CreatedAt = time.Now()
	campRound.UpdatedAt = time.Now()

	err = campRound.CreateNewCampRound()
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	fmt.Printf("%v\n", campRound)
	c.JSON(200, campRound)
	response.Response(c, errmsg.SUCCESS)
}

func GetCampRound(c *gin.Context) {

	campRounds, err := repository.GetAllCampRound()

	if err != nil {
		response.Response(c, errmsg.ERROR_ACCOUNT_NOT_EXIST)
		return
	}

	c.JSON(http.StatusOK, campRounds)
}

func GetCampRoundById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("campRoundId"))
	campRound, err := repository.GetCampRoundById(int64(id))
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	c.JSON(200, campRound)
}

func UpdateCampRound(c *gin.Context) {

}

func DeleteCampRound(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("campRoundId"))

	campRound, err := repository.DeleteCampRoundById(int64(id))

	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	c.JSON(200, campRound)
}
