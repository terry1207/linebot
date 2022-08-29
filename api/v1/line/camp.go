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
				case text_trimspace == "我要訂營地!":
					bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("img_carousel",
						&linebot.ImageCarouselTemplate{
							Columns: Add_Carousel_Imgae(),
						},
					)).Do()
				default:
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text_trimspace)).Do()
				}
			}
		}
	}

}

func Add_Carousel_Imgae() (c_i []*linebot.ImageCarouselColumn) {
	c1 := linebot.ImageCarouselColumn{
		ImageURL: "./imgae/1.jpg",
		Action: &linebot.PostbackAction{
			Label:       "A區",
			Text:        "5m*5m",
			Data:        "action=click&itemid=0",
			InputOption: linebot.InputOptionOpenRichMenu,
		},
	}

	c2 := linebot.ImageCarouselColumn{
		ImageURL: "https://example.com/bot/images/item1.jpg",
		Action: &linebot.PostbackAction{
			Label:       "B區",
			Text:        "5m*5m",
			Data:        "action=click&itemid=1",
			InputOption: linebot.InputOptionOpenRichMenu,
		},
	}
	c_i = append(c_i, &c2, &c1)
	return c_i
}

func Add_Carousel_Template() (c_t []*linebot.CarouselColumn) {

	column1 := linebot.CarouselColumn{
		ThumbnailImageURL:    "https://example.com/bot/images/item1.jpg",
		ImageBackgroundColor: "#FFFFFF",
		Title:                "A區",
		Text:                 "5m*5m",
		DefaultAction: &linebot.URIAction{
			Label: "View detail",
			URI:   "http://example.com/page/123",
		},
		Actions: []linebot.TemplateAction{

			&linebot.URIAction{
				Label: "詳細資訊",
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

	c_t = append(c_t, &column1, &column2)

	return c_t
}
