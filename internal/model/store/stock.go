package store

import (
	"time"

	"gorm.io/gorm"
)

//庫存
type Stock struct {
	gorm.Model
	Date          time.Time
	CampRoundName string `gorm:"type:varchar(100);not null;default:''" json:"CampRoundName"`
	Price         int32  `gorm:"type:smallint;not null;default:''" json:"AddressCountry"`
	TotlaNum      int32  `gorm:"type:smallint;not null;default:''" json:"TotlaNum"`
	RemainNum     int32  `gorm:"type:smallint;not null;default:''" json:"RemainNum"`
	Description   string `gorm:"type:varchar(100);default:''" json:"Description"`
}
