package repository

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Camp struct {
	gorm.Model
	CampName       string         `gorm:"type:varchar(100);not null;default:''" json:"CampName"`
	AddressCountry string         `gorm:"type:varchar(100);not null;default:''" json:"AddressCountry"`
	AddressCity    string         `gorm:"type:varchar(100);not null;default:''" json:"AddressCity"`
	AddressDetail  string         `gorm:"type:varchar(100);not null;default:''" json:"AddressDetail"`
	TagList        pq.StringArray `gorm:"type:text[]" json:"TagList"`
}

//新建營地
func (camp *Camp) CreateNewCamp() error {
	return BeginTranscation(db, func(tx *gorm.DB) error {
		return tx.Create(&camp).Error
	})

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

func DeleteCampById(Id int64) (Camp, error) {
	var camp Camp
	err := db.Where("ID=?", Id).Delete(&camp).Error

	return camp, err

}
