package order

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderRefer int
	OrderItems OrderItem `gorm:"foreignKey:OrderRefer"`
}

type OrderItem struct {
	ID   int
	Name string
	// ProductsID    int
	// PaymentAmount int
	// PaymentTotal  float32
}
