package account

import (
	"linebot/internal/config/db"

	"gorm.io/gorm"
)

//新建帳戶
func Add(account *Account) error {
	return db.BeginTranscation(db.DB, func(tx *gorm.DB) error {
		return tx.Create(&account).Error
	})
}

func GetAll() ([]Account, error) {
	var Accounts []Account
	err := db.DB.Find(&Accounts).Error

	return Accounts, err
}

func GetById(Id int64) (Account, error) {
	var GetAccount Account
	err := db.DB.Where("Id=?", Id).Find(&GetAccount).Error

	return GetAccount, err
}

func DeleteById(Id int64) (Account, error) {
	var account Account
	err := db.DB.Where("Id=?", Id).Delete(&account).Error
	return account, err
}
