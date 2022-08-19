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

func CreateTagMap(c *gin.Context) {
	var tagmap repository.TagMap
	err := c.BindJSON(&tagmap)
	if err != nil {
		return
	}

	tagmap.CreatedAt = time.Now()
	tagmap.UpdatedAt = time.Now()

	err = tagmap.CreateNewTagMap()
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	fmt.Printf("%v\n", tagmap)
	c.JSON(200, tagmap)
	response.Response(c, errmsg.SUCCESS)
}

func GetTagMap(c *gin.Context) {

	tagmaps, err := repository.GetAllTagMap()

	if err != nil {
		response.Response(c, errmsg.ERROR_ACCOUNT_NOT_EXIST)
		return
	}

	c.JSON(http.StatusOK, tagmaps)
}

func GetTagMapById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("tagmapId"))
	tagmap, err := repository.GetTagMapById(int64(id))
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	c.JSON(200, tagmap)
}

func UpdateTagMap(c *gin.Context) {

}

func DeleteTagMap(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("tagmapId"))

	tagmap, err := repository.DeleteTagMapById(int64(id))

	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	c.JSON(200, tagmap)
}
