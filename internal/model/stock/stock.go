package stock

import (
	"errors"
	"linebot/internal/config/db"
	"log"
	"time"

	"gorm.io/gorm"
)

//庫存
type Stock struct {
	gorm.Model
	Date      time.Time
	ProductId uint
	TotlaNum  int
	RemainNum int
}

func (stock Stock) Add() error {
	if stock.RemainNum > stock.TotlaNum {

		err := errors.New("remain number can't bigger than total num")
		log.Fatal(err)
	}
	return db.BeginTransaction(db.DB, func(tx *gorm.DB) error {
		return tx.Create(&stock).Error
	})
}
func GetAll() ([]Stock, error) {
	var Stocks []Stock
	err := db.DB.Find(&Stocks).Error
	return Stocks, err
}

func GetStockByDate(date time.Time) (Stock, error) {
	var stock Stock
	err := db.DB.Where("date=?", date).Find(&stock).Error
	return stock, err
}

func GetStocks_By_ID_and_DateRange(pid uint, start, end time.Time) ([]Stock, error) {

	var stocks []Stock
	err := db.DB.Where("product_id=? AND date BETWEEN ? AND ?", pid, start, end).Find(&stocks).Error
	return stocks, err
}

func UpdateStockRemain(stock Stock) error {
	return db.BeginTransaction(db.DB, func(tx *gorm.DB) error {
		return tx.Save(&stock).Error
	})
}
