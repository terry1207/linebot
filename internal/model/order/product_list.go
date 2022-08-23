package order

import (
	"time"

	"gorm.io/gorm"
)

type Product_List struct {
	gorm.Model
	ProductsID    []int
	PaymentAmount int
	PaymentTotal  float32
	Checkin       time.Time
	Checkout      time.Time
	NightGo       bool
	NightFeeTotal float32
}
