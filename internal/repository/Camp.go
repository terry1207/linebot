package repository

import (
	"gorm.io/gorm"
)

type Camp struct {
	gorm.Model
	CampName       string `gorm:"type:varchar(100);not null;default:''" json:"CampName"`
	AddressCountry string `gorm:"type:varchar(100);not null;default:''" json:"AddressCountry"`
	AddressCity    string `gorm:"type:varchar(100);not null;default:''" json:"AddressCity"`
	AddressDetail  string `gorm:"type:varchar(100);not null;default:''" json:"AddressDetail"`
	TagList        string `gorm:"type:varchar(999);not null;default:''" json:"TagList"`
}

//新建營地
func (camp *Camp) CreateNewCamp() error {
	return BeginTranscation(db, func(tx *gorm.DB) error {
		return tx.Create(&camp).Error
	})

}

func QueryCampByCampName(campName string) (*Camp, error) {
	var camp Camp
	err := db.Limit(1).Where("Name=?", campName).Find(&camp).Error
	return &camp, err

}

func GetAllCamp() ([]Camp, error) {
	var Camps []Camp
	err := db.Find(&Camps).Error

	return Camps, err

}

func GetCampById(Id int64) (*Camp, error) {
	var GetCamp Camp

	err := db.Where("ID=?", Id).Find(&GetCamp).Error

	return &GetCamp, err

}

func Delete(Id int64) (Camp, error) {
	var camp Camp
	err := db.Where("ID=?", Id).Delete(camp).Error

	return camp, err

}
