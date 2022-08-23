package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	TagName   string        `gorm:"type:varchar(10);not null;default:''" json:"TagName"`
	MapCampID pq.Int64Array `gorm:"type:text[]" json:"MapCampID"`
	TagNum    int           `gorm:"type:smallint;not null;default:0" json:"TagNum"`
}
