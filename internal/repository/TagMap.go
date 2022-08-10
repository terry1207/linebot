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

func GetAllTagMap() ([]TagMap, error) {
	var TagMaps []TagMap
	err := db.Find(&TagMaps).Error

	return TagMaps, err
}

func GetTagMapById(Id int64) (TagMap, error) {
	var GetTagMap TagMap
	err := db.Where("Id=?", Id).Find(&GetTagMap).Error

	return GetTagMap, err
}

func DeleteTagMapById(Id int64) (TagMap, error) {
	var TagMap TagMap
	err := db.Where("Id=?", Id).Delete(&TagMap).Error
	return TagMap, err
}
