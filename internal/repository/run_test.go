package repository

import (
	"fmt"
	"linebot/internal/model"
	"testing"
)

func BenchmarkCreateDataTest(b *testing.B) {

	InitDbContext()
	var (
		camp = model.Camp{

			CampName:       "哈哈營地",
			AddressCountry: "彰化",
			AddressCity:    "彰化市",
			AddressDetail:  "223號",
			TagList:        "ee",
		}

		account = model.Account{

			Password: "asdfggg",
		}

		tag = model.Tag{

			TagName: "有水池",
			TagNum:  20,
		}
		tagmap = model.TagMap{

			TagMap_CampID: "99t",
		}
	)

	CreateNewCamp(&camp)
	CreateNewUser(&account)
	CreateNewTag(&tag)
	CreateNewTagMap(&tagmap)

	var p []model.Camp
	var p1 []model.Account
	var p2 []model.Tag
	var p3 []model.TagMap

	db.Take(&p)
	db.Take(&p1)
	db.Take(&p2)
	db.Take(&p3)

	fmt.Println("camp", p)
	fmt.Println("account", p1)
	fmt.Println("tag", p2)
	fmt.Println("tagmap", p3)

	// fmt.Println("account")
	// res1 := db.Find(&model.Account{})
	// fmt.Println(res1)

	// fmt.Println("tag")
	// res2 := db.Find(&model.Tag{})
	// fmt.Println(res2)

	// fmt.Println("tagmap")
	// res3 := db.Find(&model.TagMap{})
	// fmt.Println(res3)

}

func BenchmarkPrintInt2String01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(
			"12",
		)
	}
}
