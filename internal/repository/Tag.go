package repository

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	TagName string `gorm:"type:varchar(10);not null;default:''" json:"AddressCountry"`
	TagNum  int    `gorm:"type:smallint;not null;default:0" json:"TagNum"`
}

//新建標籤
func (tag Tag) CreateNewTag() error {
	return BeginTranscation(db, func(tx *gorm.DB) error {
		return tx.Create(&tag).Error
	})
}
