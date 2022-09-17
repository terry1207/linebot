package richmenu

import (
	"fmt"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Build_RichMenu() {

	bot, err := linebot.New(

		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	Delete_All(bot)
	aid_A := "richmenu-alias-a"
	aid_B := "richmenu-alias-b"

	if CheckRichMenuAlias_exist(bot, aid_A) {
		DeleteRichMenuAlias(bot, aid_A)
	}
	if CheckRichMenuAlias_exist(bot, aid_B) {
		DeleteRichMenuAlias(bot, aid_B)
	}

	//Richmenu_Id_A := CreatRichMenu_A(bot, aid_B)

	//img_path_A := "./internal/richmenu/img/img_A.png"

	//Upload_Img(bot, Richmenu_Id_A, img_path_A)

	Richmenu_Id_B := CreatRichMenu_B(bot, aid_A)

	img_path_B := "./internal/richmenu/img/img_B.png"

	Upload_Img(bot, Richmenu_Id_B, img_path_B)

	Set_Default(bot, Richmenu_Id_B)

	//CreateRichMenuAlias(bot, aid_A, Richmenu_Id_A)
	//CreateRichMenuAlias(bot, aid_B, Richmenu_Id_B)

}

func CreatRichMenu_A(bot *linebot.Client, aid string) string {
	richMenu := linebot.RichMenu{
		Size:        linebot.RichMenuSize{Width: 2500, Height: 1686},
		Selected:    false,
		Name:        "richmenu-a",
		ChatBarText: "選單A",
		Areas: []linebot.AreaDetail{
			{
				Bounds: linebot.RichMenuBounds{X: 0, Y: 0, Width: 1200, Height: 235},
				Action: linebot.RichMenuAction{
					Type:            linebot.RichMenuActionTypeRichMenuSwitch,
					RichMenuAliasID: aid,
					Data:            "action=richmenu-changed-to-b",
				},
				// Action: linebot.RichMenuAction{
				// 	Type: linebot.RichMenuActionTypeMessage,
				// 	Text: "切換至B",
				// },
			},

			{
				Bounds: linebot.RichMenuBounds{X: 0, Y: 788, Width: 833, Height: 553},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeMessage,
					Text: "好友專屬優惠",
				},
			},

			{
				Bounds: linebot.RichMenuBounds{X: 833, Y: 788, Width: 833, Height: 553},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeMessage,
					Text: "常見問題",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{X: 1666, Y: 788, Width: 833, Height: 553},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeMessage,
					Text: "推薦我們給朋友",
				},
			},
		},
	}

	res, err := bot.CreateRichMenu(richMenu).Do()
	if err != nil {
		log.Fatal(err)
	}

	return res.RichMenuID
}

func CreatRichMenu_B(bot *linebot.Client, aid string) string {
	richMenu := linebot.RichMenu{
		Size:        linebot.RichMenuSize{Width: 2500, Height: 1686},
		Selected:    false,
		Name:        "richmenu-b",
		ChatBarText: "選單B",
		Areas: []linebot.AreaDetail{
			//{
			// 	Bounds: linebot.RichMenuBounds{X: 1251, Y: 0, Width: 1200, Height: 235},
			// 	Action: linebot.RichMenuAction{
			// 		Type:            linebot.RichMenuActionTypeRichMenuSwitch,
			// 		RichMenuAliasID: aid,
			// 		Data:            "action=richmenu-changed-to-a",
			// 	},
			// 	// Action: linebot.RichMenuAction{
			// 	// 	Type: linebot.RichMenuActionTypeMessage,
			// 	// 	Text: "切換至A",
			// 	// },
			// },
			{
				Bounds: linebot.RichMenuBounds{X: 0, Y: 234, Width: 833, Height: 553},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypePostback,
					Data: "action=search&type=go",
				},
			},

			{
				Bounds: linebot.RichMenuBounds{X: 833, Y: 235, Width: 833, Height: 553},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeURI,
					URI:  "https://www.harpersbazaar.com/tw/life/trips/g33228687/taiwan-camping-places/",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{X: 1666, Y: 235, Width: 833, Height: 553},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeMessage,
					Text: "營地介紹",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{X: 0, Y: 788, Width: 833, Height: 553},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeMessage,
					Text: "粉絲獨享",
				},
			},

			{
				Bounds: linebot.RichMenuBounds{X: 833, Y: 788, Width: 833, Height: 553},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeMessage,
					Text: "常見問題",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{X: 1666, Y: 788, Width: 833, Height: 553},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeURI,
					URI:  "tel:0909990685",
				},
			},
		},
	}

	res, err := bot.CreateRichMenu(richMenu).Do()
	if err != nil {
		log.Fatal(err)
	}

	return res.RichMenuID
}

func Upload_Img(bot *linebot.Client, id, path string) {
	if _, err := bot.UploadRichMenuImage(id, path).Do(); err != nil {
		log.Fatal(err)
	}
}

func Set_Default(bot *linebot.Client, id string) {
	if _, err := bot.SetDefaultRichMenu(id).Do(); err != nil {
		log.Fatal(err)
	}
}

func Delete_All(bot *linebot.Client) {
	res, err := bot.GetRichMenuList().Do()
	if err != nil {
		log.Fatal(err)
	}
	for _, richMenu := range res {
		if _, err := bot.DeleteRichMenu(richMenu.RichMenuID).Do(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("delete all done")
}

func CreateRichMenuAlias(bot *linebot.Client, aid, rid string) {
	if _, err := bot.CreateRichMenuAlias(aid, rid).Do(); err != nil {
		log.Fatal(err)
	}

}

func DeleteRichMenuAlias(bot *linebot.Client, aid string) {
	if _, err := bot.DeleteRichMenuAlias(aid).Do(); err != nil {
		log.Fatal("delete alias failed", err)
	}

}

func CheckRichMenuAlias_exist(bot *linebot.Client, aid string) bool {
	res, err := bot.GetRichMenuAliasList().Do()
	if err != nil {
		log.Fatal(err)
	}
	for _, alias := range res {
		if alias.RichMenuAliasID == aid {
			return true
		}
	}

	return false
}
