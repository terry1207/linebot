package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	TagName string `gorm:"type:varchar(10);not null;default:''" json:"AddressCountry"`
	TagNum  int    `gorm:"type:smallint;not null;default:0" json:"TagNum"`
}

type TagMap struct {
	gorm.Model
	TagMap_CampID string `gorm:"type:varchar(999);not null;default:''" json:"TagMapCampID"`
}
