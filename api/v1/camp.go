package v1

import (
	"linebot/internal/errmsg"
	"linebot/internal/repository"
	"linebot/internal/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCamp(c *gin.Context) {
	var test repository.Camp
	// fmt.Println(c.Query("index"))
	// fmt.Println(c.Query("name"))

	// name := c.Query("name")
	// city := c.Query("city")
	// town := c.Query("town")
	// phone := c.Query("phone")

	// test = model.Camp{

	// 	Name:        name,
	// 	City:        city,
	// 	Town:        town,
	// 	PhoneNumber: phone,
	// }
	// fmt.Println(test)

	err := test.CreateNewCamp()
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

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
