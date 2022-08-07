package repository

import (
	"fmt"
	"log"
	"testing"
)

func BenchmarkCreateDataTest(b *testing.B) {

	InitDbContext()
	var (
		camp = Camp{

			CampName:       "哈哈營地",
			AddressCountry: "彰化",
			AddressCity:    "彰化市",
			AddressDetail:  "223號",
			TagList:        "ee",
		}
		camp2 = Camp{

			CampName:       "哈哈營地",
			AddressCountry: "彰化",
			AddressCity:    "彰化市",
			AddressDetail:  "223號",
			TagList:        "ee",
		}

		account = Account{

			Password: "asdfggg",
		}

		tag = Tag{

			TagName: "有水池",
			TagNum:  20,
		}
		tagmap = TagMap{

			TagMap_CampID: "99t",
		}
	)

	camp.CreateNewCamp()
	camp2.CreateNewCamp()

	account.CreateNewUser()
	tag.CreateNewTag()
	tagmap.CreateNewTagMap()

	getAllcamp, err := GetAllCamp()
	if err != nil {
		log.Println(err)
	}
	for _, r := range getAllcamp {
		fmt.Println(r)
	}

}
