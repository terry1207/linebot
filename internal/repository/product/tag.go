package repository

import (
	"fmt"
	"linebot/internal/model/product"
	"linebot/internal/repository/db"
	"linebot/pkg/tool"

	"gorm.io/gorm"
)

//新建標籤
func CreateNewTag(tag *product.Tag) error {
	return db.BeginTranscation(db.DB, func(tx *gorm.DB) error {
		return tx.Create(&tag).Error
	})
}

func GetAllTag() ([]product.Tag, error) {
	var Tags []product.Tag
	err := db.DB.Find(&Tags).Error

	return Tags, err
}

func GetTagById(Id int64) (product.Tag, error) {
	var GetTag product.Tag
	err := db.DB.Where("Id=?", Id).Find(&GetTag).Error

	return GetTag, err
}

func DeleteTagById(Id int64) (product.Tag, error) {
	var tag product.Tag
	err := db.DB.Where("Id=?", Id).Delete(&tag).Error
	return tag, err
}

func UpdateTag_from_CampId(tagId, campId int64) error {
	tag, err := GetTagById(tagId)
	if err != nil {
		fmt.Printf("get tag by id failed\n %s", err.Error())

	}

	if !tool.IsExist_in_Arr(campId, tag.MapCampID) {
		tag.MapCampID = append(tag.MapCampID, int64(campId))
		tag.TagNum++
		return db.BeginTranscation(db.DB, func(tx *gorm.DB) error {
			return tx.Save(&tag).Error
		})
	} else {
		fmt.Printf("CampId exist Please CheckOut\n")
	}

	return nil
}
