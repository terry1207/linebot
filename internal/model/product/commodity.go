package product

import "gorm.io/gorm"

//營地商品
type CampRound struct {
	gorm.Model
	CampRoundName string `gorm:"type:varchar(100);not null;default:''" json:"CampRoundName"`
	Price         int32  `gorm:"type:smallint;not null;default:''" json:"AddressCountry"`
	Description   string `gorm:"type:varchar(100);default:''" json:"Description"`
}
