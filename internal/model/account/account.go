package account

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name     string
	Password string `gorm:"type:varchar(100);not null;default:''" json:"password"`
}
