package model

import "gorm.io/gorm"

type Camp struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;default:''" json:"name"`
	City string `gorm:"type:varchar(20);not null;default:''" json:"city"`
	Town string `gorm:"type:varchar(20);not null;default:''" json:"town"`
}
