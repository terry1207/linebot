package v1

import (
	"linebot/internal/errmsg"
	"linebot/internal/model"
	"linebot/internal/repository"
	"linebot/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCampInfo(c *gin.Context) {
	var test model.Camp

	test = model.Camp{
		Name: "哈囉營地",
		City: "南投",
		Town: "埔里",
	}

	err := repository.CreateNewCamp(&test)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	response.Response(c, errmsg.SUCCESS)
}

func GetCampInfo(c *gin.Context) {
	var query = "哈囉營地"

	camp, err := repository.QueryCampByCampName(query)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	if camp.Name == "" {
		response.Response(c, errmsg.ERROR_ACCOUNT_NOT_EXIST)
		return
	}

	c.JSON(http.StatusOK, camp)
}
