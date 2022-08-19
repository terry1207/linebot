package repository

import (
	"fmt"
	"linebot/pkg/tool"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	TagName   string        `gorm:"type:varchar(10);not null;default:''" json:"TagName"`
	MapCampID pq.Int64Array `gorm:"type:text[]" json:"MapCampID"`
	TagNum    int           `gorm:"type:smallint;not null;default:0" json:"TagNum"`
}

//新建標籤
func (tag Tag) CreateNewTag() error {
	return BeginTranscation(db, func(tx *gorm.DB) error {
		return tx.Create(&tag).Error
	})
}

func GetAllTag() ([]Tag, error) {
	var Tags []Tag
	err := db.Find(&Tags).Error

	return Tags, err
}

func GetTagById(Id int64) (Tag, error) {
	var GetTag Tag
	err := db.Where("Id=?", Id).Find(&GetTag).Error

	return GetTag, err
}

func DeleteTagById(Id int64) (Tag, error) {
	var Tag Tag
	err := db.Where("Id=?", Id).Delete(&Tag).Error
	return Tag, err
}

func UpdateTag_from_CampId(tagId, campId int64) error {
	tag, err := GetTagById(tagId)
	if err != nil {
		fmt.Printf("get tag by id failed\n %s", err.Error())

	}

	if !tool.IsExist_in_Arr(campId, tag.MapCampID) {
		tag.MapCampID = append(tag.MapCampID, int64(campId))
		tag.TagNum++
		return BeginTranscation(db, func(tx *gorm.DB) error {
			return tx.Save(&tag).Error
		})
	} else {
		fmt.Printf("CampId exist Please CheckOut\n")
	}

	return nil
}
