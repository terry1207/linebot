package account

import (
	"linebot/internal/model/account"
	"linebot/internal/repository/db"

	"gorm.io/gorm"
)

//新建帳戶
func AddAccount(account *account.Account) error {
	return db.BeginTranscation(db.DB, func(tx *gorm.DB) error {
		return tx.Create(&account).Error
	})
}

func GetAllAccount() ([]account.Account, error) {
	var Accounts []account.Account
	err := db.DB.Find(&Accounts).Error

	return Accounts, err
}

func GetAccountById(Id int64) (account.Account, error) {
	var GetAccount account.Account
	err := db.DB.Where("Id=?", Id).Find(&GetAccount).Error

	return GetAccount, err
}

func DeleteAccountById(Id int64) (account.Account, error) {
	var account account.Account
	err := db.DB.Where("Id=?", Id).Delete(&account).Error
	return account, err
}
