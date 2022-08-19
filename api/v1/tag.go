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

func CreateTag(c *gin.Context) {
	var tag repository.Tag
	err := c.BindJSON(&tag)
	if err != nil {
		return
	}

	tag.CreatedAt = time.Now()
	tag.UpdatedAt = time.Now()

	err = tag.CreateNewTag()
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	fmt.Printf("%v\n", tag)
	c.JSON(200, tag)
	response.Response(c, errmsg.SUCCESS)
}

func GetTag(c *gin.Context) {

	tags, err := repository.GetAllTag()

	if err != nil {
		response.Response(c, errmsg.ERROR_ACCOUNT_NOT_EXIST)
		return
	}

	c.JSON(http.StatusOK, tags)
}

func GetTagById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("tagId"))
	tag, err := repository.GetTagById(int64(id))
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	c.JSON(200, tag)
}

func UpdateTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("tagId"))

	tag, err := repository.DeleteTagById(int64(id))

	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	c.JSON(200, tag)
}
