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

func QueryCampByCampName(campName string) (model.Camp, error) {
	var camp model.Camp
	err := db.Limit(1).Where("Name=?", campName).Find(&camp).Error
	return camp, err

}
