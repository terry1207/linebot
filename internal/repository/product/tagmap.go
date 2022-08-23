package repository

import (
	"linebot/internal/model/product"
	"linebot/internal/repository/db"

	"gorm.io/gorm"
)

//新建標籤字典
func CreateNewTagMap(tagmap *product.TagMap) error {
	return db.BeginTranscation(db.DB, func(tx *gorm.DB) error {
		return tx.Create(&tagmap).Error
	})
}

func GetAllTagMap() ([]product.TagMap, error) {
	var TagMaps []product.TagMap
	err := db.DB.Find(&TagMaps).Error

	return TagMaps, err
}

func GetTagMapById(Id int64) (product.TagMap, error) {
	var GetTagMap product.TagMap
	err := db.DB.Where("Id=?", Id).Find(&GetTagMap).Error

	return GetTagMap, err
}

func DeleteTagMapById(Id int64) (product.TagMap, error) {
	var tagMap product.TagMap
	err := db.DB.Where("Id=?", Id).Delete(tagMap).Error
	return tagMap, err
}
