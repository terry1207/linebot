package route

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	//"linebot/api/v1/front"

	"linebot/internal/repository"

	"linebot/internal/user_search"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func RepeatHandler(c *gin.Context) {
	var r = 10
	var buffer bytes.Buffer
	for i := 0; i < r; i++ {
		buffer.WriteString("Hello from Go!\n")
	}
	c.String(http.StatusOK, buffer.String())

}

func DbTest(c *gin.Context) {
	repository.InsertTest()
}

var post bool
var bot *linebot.Client
var err error

func init() {
	bot, err = linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func ReplyMessage(c *gin.Context) {

	events, err := bot.ParseRequest(c.Request)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Writer.WriteHeader(400)
		} else {
			c.Writer.WriteHeader(500)
		}
		return
	}

	for _, event := range events {

		if event.Postback.Data != "" {
			fmt.Println(event.Postback.Data)
		}

		var id string
		switch {
		case event.Source.UserID != "":
			id = event.Source.UserID
		case event.Source.GroupID != "":
			id = event.Source.GroupID
		}

		if user_search.ID_Search[id] == nil {
			user_search.ID_Search[id] = &user_search.Search_Weath{}
			fmt.Println("Init id", user_search.ID_Search[id])
		}

		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				text_trimspace := strings.TrimSpace(message.Text)

				switch {
				case text_trimspace == "e":
					fmt.Println(user_search.ID_Search)
				case user_search.ID_Search[id].Start_Search:
					user_search.ID_Search[id].Start_Search = false
					user_search.ID_Search[id].Search_Region = true
					user_search.ID_Search[id].Search_Input.RegionName = text_trimspace

					city := user_search.ID_Search[id].Quick_Reply_City()

					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("選擇縣市").WithQuickReplies(&city)).Do()
				case user_search.ID_Search[id].Search_Region:
					user_search.ID_Search[id].Search_Region = false
					user_search.ID_Search[id].Search_City = true
					user_search.ID_Search[id].Search_Input.Location = text_trimspace

					towns := user_search.ID_Search[id].Quick_Reply_Town()
					fmt.Println(towns)
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("選擇鄉鎮區域").WithQuickReplies(&towns)).Do()

				case user_search.ID_Search[id].Search_City:
					switch text_trimspace {
					case string(user_search.Previous_Page):
						user_search.ID_Search[id].Search_Town_Index--

						towns := user_search.ID_Search[id].Quick_Reply_Town()
						fmt.Println(towns)
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("選擇鄉鎮區域").WithQuickReplies(&towns)).Do()

					case string(user_search.Next_Page):
						user_search.ID_Search[id].Search_Town_Index++

						towns := user_search.ID_Search[id].Quick_Reply_Town()
						fmt.Println(towns)
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("選擇鄉鎮區域").WithQuickReplies(&towns)).Do()

					default:
						user_search.ID_Search[id].Search_City = false
						user_search.ID_Search[id].Search_Input.LocationName = text_trimspace

						fmt.Println("call 前資訊", user_search.ID_Search)

						result := user_search.Call_Weath_Api(&user_search.ID_Search[id].Search_Input)

						delete(user_search.ID_Search, id)
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(result)).Do()
					}

				default:
					switch text_trimspace {

					case "查詢降雨機率":
						region := user_search.Quick_Reply_Region()
						user_search.ID_Search[id].Start_Search = true

						fmt.Println("user_search.ID_Search Init:", &user_search.ID_Search)
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("選擇區域").WithQuickReplies(&region)).Do()
					case "t":
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("replytoken:"+event.ReplyToken)).Do()
					case "w":
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("WebhookEventID:"+event.WebhookEventID)).Do()
					case "g":
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Source.GroupID:"+event.Source.GroupID)).Do()
					case "r":
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Source.RoomID:"+event.Source.RoomID)).Do()
					case "u":
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Source.UserID:"+event.Source.UserID)).Do()
					case "p":

						bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("this is a button template", &linebot.ButtonsTemplate{
							ThumbnailImageURL:    "https://example.com/bot/images/image.jpg",
							ImageAspectRatio:     "rectangle",
							ImageSize:            "cover",
							ImageBackgroundColor: "#FFFFFF",
							Title:                "Menu",
							Text:                 "Please select",
							DefaultAction: &linebot.URIAction{
								Label: "View detail",
								URI:   "http://example.com/page/123",
							},
							Actions: []linebot.TemplateAction{
								&linebot.PostbackAction{
									Label: "Buy",
									Data:  "action=buy&itemid=123",
									Text:  "收到囉",
								},

								&linebot.URIAction{
									Label: "View detail",
									URI:   "http://example.com/page/123",
								},
							},
						})).Do()
					case "d":
						bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("time", &linebot.ButtonsTemplate{
							Title: "Menu",
							Text:  "Please select",
							DefaultAction: &linebot.DatetimePickerAction{
								Label: "選擇時間",
								Mode:  "date",
								Data:  "type=1",
							},
							Actions: []linebot.TemplateAction{
								&linebot.DatetimePickerAction{
									Label: "選擇時間",
									Mode:  "date",
									Data:  "type=1",
								},
							},
						},
						)).Do()
					case "收到囉":
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(event.Postback.Data)).Do()
					default:
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("請問要查詢什麼？").WithQuickReplies(&linebot.QuickReplyItems{
							Items: []*linebot.QuickReplyButton{
								{
									Action: &linebot.MessageAction{
										Label: "查詢降雨機率",
										Text:  "查詢降雨機率",
									},
								},
							},
						})).Do()

					}

				}
			}
		}

	}

}
