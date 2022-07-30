package repository

import (
	"linebot/internal/model"

	"gorm.io/gorm"
)

//新建營地
func CreateNewCamp(camp *model.Camp) error {
	return BeginTranscation(db, func(tx *gorm.DB) error {
		return tx.Create(&camp).Error
	})

}

func CreateNewUser(account *model.Account) error {
	return BeginTranscation(db, func(tx *gorm.DB) error {
		return tx.Create(&account).Error
	})
}

func CreateNewTag(tag *model.Tag) error {
	return BeginTranscation(db, func(tx *gorm.DB) error {
		return tx.Create(&tag).Error
	})
}

func CreateNewTagMap(tagmap *model.TagMap) error {
	return BeginTranscation(db, func(tx *gorm.DB) error {
		return tx.Create(&tagmap).Error
	})
}

func QueryCampByCampName(campName string) (model.Camp, error) {
	var camp model.Camp
	err := db.Limit(1).Where("Name=?", campName).Find(&camp).Error
	return camp, err

}
