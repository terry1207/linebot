package repository

import (
	"linebot/internal/model/product"
	"linebot/internal/repository/db"

	"gorm.io/gorm"
)

//新建營地商品
func CreateNewCampRound(commodity *product.CampRound) error {
	return db.BeginTranscation(db.DB, func(tx *gorm.DB) error {
		return tx.Create(&commodity).Error
	})

}

func GetAllCampRound() ([]product.CampRound, error) {
	var CampRounds []product.CampRound
	err := db.DB.Find(&CampRounds).Error

	return CampRounds, err

}

func GetCampRoundById(Id int64) (*product.CampRound, error) {
	var GetCampRound product.CampRound

	err := db.DB.Where("ID=?", Id).Find(&GetCampRound).Error

	return &GetCampRound, err

}

func DeleteCampRoundById(Id int64) (product.CampRound, error) {
	var commodity product.CampRound
	err := db.DB.Where("ID=?", Id).Delete(&commodity).Error

	return commodity, err

}
