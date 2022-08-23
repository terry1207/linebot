package product

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Date        time.Time `gorm:"not null"`
	Name        string    `gorm:"not null"`
	Price       float32   `gorm:"not null;default:0"`
	NightFee    float32   `gorm:"not null;default:0"`
	Size        string
	Image       int
	Description string
}
