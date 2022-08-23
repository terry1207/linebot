package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TagMap struct {
	gorm.Model
	TagMap_CampID string `gorm:"type:varchar(999);not null;default:''" json:"TagMapCampID"`
	TagMap_Camp   pq.StringArray
}
