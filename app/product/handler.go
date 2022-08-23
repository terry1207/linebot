package product

import (
	"linebot/internal/model/product"
	"linebot/internal/repository/db"

	"gorm.io/gorm"
)

//新建帳戶
func Add(product *product.Product) error {
	return db.BeginTranscation(db.DB, func(tx *gorm.DB) error {
		return tx.Create(&product).Error
	})
}

func GetAll() ([]product.Product, error) {
	var Products []product.Product
	err := db.DB.Find(&Products).Error

	return Products, err
}

func GetById(Id int64) (product.Product, error) {
	var GetProduct product.Product
	err := db.DB.Where("Id=?", Id).Find(&GetProduct).Error

	return GetProduct, err
}

func DeleteById(Id int64) (product.Product, error) {
	var product product.Product
	err := db.DB.Where("Id=?", Id).Delete(&product).Error
	return product, err
}
