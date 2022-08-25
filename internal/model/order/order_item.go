package order

import (
	"time"

	"gorm.io/gorm"
)

//訂購營地帳數 EX A區 3帳
type OrderItem struct {
	gorm.Model
	//訂單號碼 table orders(id)
	OrderID      int
	ProductId    int
	Amount       int
	PaymentTotal float32
	Checkin      time.Time
	Checkout     time.Time
}
