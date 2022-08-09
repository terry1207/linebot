package repository

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	InitDbContext()
}

func BenchmarkCampTest(b *testing.B) {

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
	)

	camp.CreateNewCamp()
	camp2.CreateNewCamp()

	getAllcamp, err := GetAllCamp()
	if err != nil {
		log.Println(err)
	}
	for _, r := range getAllcamp {
		fmt.Println(r)
	}

	getcampbyid, err := GetCampById(1)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(getcampbyid)

	delcampbyid, err := DeleteCampById(1)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(delcampbyid)

}

func BenchmarkUserTest(b *testing.B) {

	var (
		account = User{

			Password: "asdfggg",
		}

		account2 = User{

			Password: "asdfggg",
		}
	)

	account.CreateNewUser()
	account2.CreateNewUser()

	getAlluser, err := GetAllUser()
	if err != nil {
		log.Println(err)
	}
	for _, r := range getAlluser {
		fmt.Println(r)
	}

	getuserbyid, err := GetUserById(1)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(getuserbyid)

	deluserbyid, err := DeleteUserById(1)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(deluserbyid)

}

func BenchmarkTagTest(b *testing.B) {

	var (
		tag = Tag{

			TagName: "有水池",
			TagNum:  20,
		}

		tag2 = Tag{

			TagName: "有水池",
			TagNum:  20,
		}
	)

	tag.CreateNewTag()
	tag2.CreateNewTag()

	getAlltag, err := GetAllTag()
	if err != nil {
		log.Println(err)
	}
	for _, r := range getAlltag {
		fmt.Println(r)
	}

	getTagbyid, err := GetTagById(1)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(getTagbyid)

	delTagbyid, err := DeleteTagById(1)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(delTagbyid)

}

func BenchmarkTagMapTest(b *testing.B) {

	var (
		tagmap = TagMap{

			TagMap_CampID: "99t",
		}

		tagmap2 = TagMap{

			TagMap_CampID: "99t",
		}
	)

	tagmap.CreateNewTagMap()
	tagmap2.CreateNewTagMap()

	getAlltagmap, err := GetAllTagMap()
	if err != nil {
		log.Println(err)
	}
	for _, r := range getAlltagmap {
		fmt.Println(r)
	}

	getTagMapbyid, err := GetTagMapById(1)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(getTagMapbyid)

	delTagMapbyid, err := DeleteTagMapById(1)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(delTagMapbyid)

}
