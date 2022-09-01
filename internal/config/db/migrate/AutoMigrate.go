package migrate

import (
	"linebot/internal/config/db"
	"linebot/internal/model/account"
	"linebot/internal/model/order"
	"linebot/internal/model/product"
)

func Init() {
	DropTable(&order.Order{}, &order.OrderItem{}, &product.Product{}, &product.Price{}, &account.Account{})
	// // AutoMigrate(&product.Product{})
	// DB.Migrator().CreateConstraint(&order.Order{}, "OrderItem")
	AutoMigrate(&order.Order{}, &order.OrderItem{})
	AutoMigrate(&product.Product{}, &product.Price{})
	AutoMigrate(&account.Account{})
}

// AutoMigrate run auto migration
func AutoMigrate(values ...interface{}) {
	for _, value := range values {

		db.DB.AutoMigrate(value)
	}
}

func DropTable(values ...interface{}) {
	for _, value := range values {
		db.DB.Migrator().DropTable(value)
	}
}
