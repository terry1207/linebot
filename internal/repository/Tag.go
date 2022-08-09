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

func GetAllTag() ([]Tag, error) {
	var Tags []Tag
	err := db.Find(&Tags).Error

	return Tags, err
}

func GetTagById(Id int64) (Tag, error) {
	var GetTag Tag
	err := db.Where("Id=?", Id).Find(&GetTag).Error

	return GetTag, err
}

func DeleteTagById(Id int64) (Tag, error) {
	var Tag Tag
	err := db.Where("Id=?", Id).Delete(&Tag).Error
	return Tag, err
}
