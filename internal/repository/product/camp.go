package repository

import (
	"linebot/internal/model/product"
	"linebot/internal/repository/db"

	"gorm.io/gorm"
)

//新建營地
func CreateNewCamp(camp *product.Camp) error {
	return db.BeginTranscation(db.DB, func(tx *gorm.DB) error {
		return tx.Create(&camp).Error
	})

}

func GetAllCamp() ([]product.Camp, error) {
	var Camps []product.Camp
	err := db.DB.Find(&Camps).Error

	return Camps, err

}

func GetCampById(Id int64) (*product.Camp, error) {
	var GetCamp product.Camp

	err := db.DB.Where("ID=?", Id).Find(&GetCamp).Error

	return &GetCamp, err

}

func DeleteCampById(Id int64) (product.Camp, error) {
	var camp product.Camp
	err := db.DB.Where("ID=?", Id).Delete(&camp).Error

	return camp, err

}
