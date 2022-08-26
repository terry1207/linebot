package line

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func CampReply(c *gin.Context) {
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

		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				text_trimspace := strings.TrimSpace(message.Text)

				switch {
				case text_trimspace == "我要訂營地！":

					column1 := linebot.CarouselColumn{
						ThumbnailImageURL:    "https://example.com/bot/images/item1.jpg",
						ImageBackgroundColor: "#FFFFFF",
						Title:                "this is menu",
						Text:                 "description",
						DefaultAction: &linebot.URIAction{
							Label: "View detail",
							URI:   "http://example.com/page/123",
						},
						Actions: []linebot.TemplateAction{
							&linebot.PostbackAction{
								Label: "Buy",
								Data:  "action=buy&itemid=111",
							},
							&linebot.PostbackAction{
								Label: "Add to chart",
								Data:  "action=buy&itemid=111",
							},
							&linebot.URIAction{
								Label: "View detail",
								URI:   "http://example.com/page/111",
							},
						},
					}
					column2 := linebot.CarouselColumn{
						ThumbnailImageURL:    "https://example.com/bot/images/item2.jpg",
						ImageBackgroundColor: "#000000",
						Title:                "this is menu",
						Text:                 "description",
						DefaultAction: &linebot.URIAction{
							Label: "View detail",
							URI:   "http://example.com/page/222",
						},
						Actions: []linebot.TemplateAction{
							&linebot.PostbackAction{
								Label: "Buy",
								Data:  "action=buy&itemid=222",
							},
							&linebot.PostbackAction{
								Label: "Add to chart",
								Data:  "action=buy&itemid=222",
							},
							&linebot.URIAction{
								Label: "View detail",
								URI:   "http://example.com/page/222",
							},
						},
					}
					bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("carousel template",
						&linebot.CarouselTemplate{
							Columns:          []*linebot.CarouselColumn{&column1, &column2},
							ImageAspectRatio: "rectangle",
							ImageSize:        "cover",
						})).Do()
				default:
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text_trimspace)).Do()
				}
			}
		}
	}

}
