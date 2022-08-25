package order

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderSN     int    `gorm:"comment:訂單序號"`
	UserName    string `gorm:"comment:登記者名字"`
	PhoneNumber string `gorm:"comment:登記者電話"`
	Email       string `gorm:"comment:登記者email"`
}
