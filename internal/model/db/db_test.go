package db

import (
	"fmt"
	"linebot/internal/model/order"
	"testing"
)

func TestDB(t *testing.T) {
	InitDbContext()
	DropTable(&order.OrderItem{}, &order.Order{})
	// AutoMigrate(&product.Product{})
	AutoMigrate(&order.OrderItem{}, &order.Order{})

	fmt.Println(DB.Migrator().HasTable(&order.OrderItem{}))
	fmt.Println(DB.Migrator().HasTable(&order.Order{}))
	DB.Create(&order.OrderItem{Name: "XYY"})
	DB.Create(&order.Order{
		OrderItems: order.OrderItem{
			ID:   1,
			Name: "XYZ",
		},
	})

	var r order.Order
	DB.Find(&r)
	fmt.Println(r.OrderItems)

	var t1 order.OrderItem
	DB.Find(&t1)
	fmt.Println(t1)
}
