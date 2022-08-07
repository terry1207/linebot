package repository

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Password string `gorm:"type:varchar(100);not null;default:''" json:"password"`
}

//新建帳戶
func (account *Account) CreateNewUser() error {
	return BeginTranscation(db, func(tx *gorm.DB) error {
		return tx.Create(&account).Error
	})
}
