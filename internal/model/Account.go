package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Password string `gorm:"type:varchar(100);not null;default:''" json:"password"`
}
