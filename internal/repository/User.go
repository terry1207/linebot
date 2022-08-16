package repository

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password string `gorm:"type:varchar(100);not null;default:''" json:"password"`
}

//新建帳戶
func (account *User) CreateNewUser() error {
	return BeginTranscation(db, func(tx *gorm.DB) error {
		return tx.Create(&account).Error
	})
}

func GetAllUser() ([]User, error) {
	var Users []User
	err := db.Find(&Users).Error

	return Users, err
}

func GetUserById(Id int64) (User, error) {
	var GetUser User
	err := db.Where("Id=?", Id).Find(&GetUser).Error

	return GetUser, err
}

func DeleteUserById(Id int64) (User, error) {
	var user User
	err := db.Where("Id=?", Id).Delete(&user).Error
	return user, err
}
