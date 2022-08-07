package repository

import (
	"gorm.io/gorm"
)

type TagMap struct {
	gorm.Model
	TagMap_CampID string `gorm:"type:varchar(999);not null;default:''" json:"TagMapCampID"`
}

//新建標籤字典
func (tagmap TagMap) CreateNewTagMap() error {
	return BeginTranscation(db, func(tx *gorm.DB) error {
		return tx.Create(&tagmap).Error
	})
}
