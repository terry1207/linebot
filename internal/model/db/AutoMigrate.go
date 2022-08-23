package db

import (
	"linebot/internal/model/order"
)

func Init() {
	//DropTable(&product.Product{})
	DropTable(&order.Product_List{})
	// AutoMigrate(&product.Product{})
	AutoMigrate(&order.Product_List{})

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
