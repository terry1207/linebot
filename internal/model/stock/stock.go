package stock

import (
	"linebot/internal/config/db"
	"time"

	"gorm.io/gorm"
)

//庫存
type Stock struct {
	gorm.Model
	Date      time.Time
	ProductId int
	TotlaNum  int32
	RemainNum int32
}

func (stock Stock) Add() error {
	return db.BeginTransaction(db.DB, func(tx *gorm.DB) error {
		return tx.Create(&stock).Error
	})
}

func GetStockByDate(date time.Time) (Stock, error) {
	var stock Stock
	err := db.DB.Where("date=?", date).Find(&stock).Error
	return stock, err
}

func GetStocks_By_ID_and_DateRange(pid int, start, end time.Time) ([]Stock, error) {

	var stocks []Stock
	err := db.DB.Where("product_id=? AND date BETWEEN ? AND ?", pid, start, end).Find(&stocks).Error
	return stocks, err
}

func UpdateStockRemain(stock Stock) error {
	return db.BeginTransaction(db.DB, func(tx *gorm.DB) error {
		return tx.Save(&stock).Error
	})
}
