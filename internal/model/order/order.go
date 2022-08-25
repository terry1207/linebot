package order

<<<<<<< HEAD
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
=======
import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderSN     int    `gorm:"comment:訂單序號"`
	UserName    string `gorm:"comment:登記者名字"`
	PhoneNumber string `gorm:"comment:登記者電話"`
	Email       string `gorm:"comment:登記者email"`
>>>>>>> refs/remotes/develop/develop
}
