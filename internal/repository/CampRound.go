package repository

import (
	"time"

	"gorm.io/gorm"
)

type CampRound struct {
	gorm.Model
	Date      time.Time
	RoundName string `gorm:"type:varchar(100);not null;default:''" json:"RoundName"`
	TotlaNum  int32  `gorm:"type:varchar(100);not null;default:''" json:"TotlaNum"`
	RemainNum int32  `gorm:"type:varchar(100);not null;default:''" json:"RemainNum"`
	Price     int32  `gorm:"type:smallint;not null;default:0" json:"Price"`
}

//新建營地庫存
func (campRound *CampRound) CreateNewCampRound() error {
	return BeginTranscation(db, func(tx *gorm.DB) error {
		return tx.Create(&campRound).Error
	})

}

func GetAllCampRound() ([]CampRound, error) {
	var CampRounds []CampRound
	err := db.Find(&CampRounds).Error

	return CampRounds, err

}

func GetCampRoundById(Id int64) (*CampRound, error) {
	var GetCampRound CampRound

	err := db.Where("ID=?", Id).Find(&GetCampRound).Error

	return &GetCampRound, err

}

func DeleteCampRoundById(Id int64) (CampRound, error) {
	var campRound CampRound
	err := db.Where("ID=?", Id).Delete(&campRound).Error

	return campRound, err

}
