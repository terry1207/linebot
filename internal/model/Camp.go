package model

import "gorm.io/gorm"

type Camp struct {
	gorm.Model
	CampName       string `gorm:"type:varchar(100);not null;default:''" json:"CampName"`
	AddressCountry string `gorm:"type:varchar(10);not null;default:''" json:"AddressCountry"`
	AddressCity    string `gorm:"type:varchar(10);not null;default:''" json:"AddressCity"`
	AddressDetail  string `gorm:"type:varchar(100);not null;default:''" json:"AddressDetail"`
	//TagList        []string `gorm:"type:varchar(999);not null;default:''" json:"TagList"`
}
