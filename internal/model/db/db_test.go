package db

import (
	"fmt"
	"linebot/internal/model/order"
	"testing"
)

func TestDB(t *testing.T) {
	InitDbContext()
	Init()

	var test order.Product_List

	DB.Create(&test)
	var tests []order.Product_List
	DB.Find(&tests)
	fmt.Println(tests)

}
