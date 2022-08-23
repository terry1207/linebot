package repository

import (
	"linebot/internal/model/store"
	"linebot/internal/repository/db"

	"gorm.io/gorm"
)

//新建營地庫存
func CreateNewStock(stock *store.Stock) error {
	return db.BeginTranscation(db.DB, func(tx *gorm.DB) error {
		return tx.Create(&stock).Error
	})

}

func GetAllStock() ([]store.Stock, error) {
	var Stocks []store.Stock
	err := db.DB.Find(&Stocks).Error

	return Stocks, err

}

func GetStockById(Id int64) (*store.Stock, error) {
	var GetStock store.Stock

	err := db.DB.Where("ID=?", Id).Find(&GetStock).Error

	return &GetStock, err

}

func DeleteStockById(Id int64) (store.Stock, error) {
	var stock store.Stock
	err := db.DB.Where("ID=?", Id).Delete(&stock).Error

	return stock, err

}
