package db

import (
	"linebot/internal/model/order"
)

func Init() {
	//DropTable(&product.Product{})
	DropTable(&order.OrderItem{}, &order.Order{})
	// AutoMigrate(&product.Product{})
	AutoMigrate(&order.OrderItem{}, &order.Order{})

}

// AutoMigrate run auto migration
func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		DB.AutoMigrate(value)
	}
}

func DropTable(values ...interface{}) {
	for _, value := range values {
		DB.Migrator().DropTable(value)
	}
}
