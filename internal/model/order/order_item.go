package order

import (
	"linebot/internal/config/db"
	"time"

	"gorm.io/gorm"
)

//訂購營地帳數 EX A區 3帳
type OrderItem struct {
	gorm.Model
	//訂單號碼 table orders(id)
	OrderSN      string
	ProductId    int
	Amount       int
	PaymentTotal float32
	Checkin      time.Time
	Checkout     time.Time
}

func (orderitem *OrderItem) Add() error {
	return db.BeginTranscation(db.DB, func(tx *gorm.DB) error {
		return tx.Create(&orderitem).Error
	})
}

func GetOrderItemByOrderSN(order_sn string) (OrderItem, error) {
	var GetOrderItem OrderItem
	err := db.DB.Where("order_sn=?", order_sn).Find(&GetOrderItem).Error
	return GetOrderItem, err
}

func DeleteById(Id int) error {
	var orderitem OrderItem
	return db.DB.Where("Id=?", Id).Delete(&orderitem).Error

}
