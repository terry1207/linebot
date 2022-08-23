package product

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
