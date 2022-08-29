package product

import (
	"fmt"
	"linebot/internal/errmsg"
	"linebot/internal/model/product"

	"linebot/internal/response"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product product.Product
	err := c.BindJSON(&product)
	if err != nil {
		return
	}

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	err = Add(&product)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	fmt.Printf("%v\n", product)
	c.JSON(200, product)
	response.Response(c, errmsg.SUCCESS)
}

func GetProduct(c *gin.Context) {

	products, err := GetAll()

	if err != nil {
		response.Response(c, errmsg.ERROR_ACCOUNT_NOT_EXIST)
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("ProductId"))
	product, err := GetById(int64(id))
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	c.JSON(200, product)
}

func UpdateProduct(c *gin.Context) {

}

func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("productId"))

	product, err := DeleteById(int64(id))

	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	c.JSON(200, product)
}

func UploadImage(c *gin.Context) {

	file, _ := c.FormFile("file")

	//filename := file.Filename
	filename := "test.png"
	if err := c.SaveUploadedFile(file, "./image/"+filename); err != nil {

		//自己完成信息提示
		return
	}
	c.String(200, "Success")
}
