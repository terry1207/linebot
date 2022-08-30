package line

import (
	"fmt"
	"linebot/internal/model/product"
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
				// if product, ok := Is_Name_Exist(text_trimspace); ok {
				// 	bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("img carousel",
				// 		&linebot.ImageCarouselTemplate{
				// 			Columns: Img_Carousel_CampRound_Info(product),
				// 		}))
				// }
				switch {

				case text_trimspace == "我要訂營地!":
					bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("img carousel",
						&linebot.ImageCarouselTemplate{
							Columns: Add_Carousel_Imgae(),
						},
					)).Do()
				case text_trimspace == "營地介紹":
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("選擇分區").WithQuickReplies(Quick_Reply_CampRoundName())).Do()
				case text_trimspace == "car":
					bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("carousel template",
						&linebot.CarouselTemplate{
							Columns:          Add_Carousel_Template(),
							ImageAspectRatio: "rectangle",
							ImageSize:        "cover",
						},
					)).Do()
				default:
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text_trimspace)).Do()
				}
			}
		}
	}

}

//快速回覆營位分區名稱
func Quick_Reply_CampRoundName() (q_p *linebot.QuickReplyItems) {
	products, _ := product.GetAll()
	for _, p := range products {
		tmp := &linebot.QuickReplyButton{
			Action: &linebot.MessageAction{
				Label: p.CampRoundName,
				Text:  p.CampRoundName,
			},
		}
		q_p.Items = append(q_p.Items, tmp)
	}
	fmt.Println("Quick_Reply_CampRoundName", q_p)
	return q_p
}

//確認輸入營位分區名是否存在
func Is_Name_Exist(name string) (product.Product, bool) {
	fmt.Println("Is_Name_Exist 輸入", name)
	products, _ := product.GetAll()
	var tmp product.Product
	for _, p := range products {
		if p.CampRoundName == name {
			tmp = p
			fmt.Println("名稱存在")
			return tmp, true
		}
	}
	return tmp, false
}
func Add_Carousel_Imgae() (c_i []*linebot.ImageCarouselColumn) {
	c1 := linebot.ImageCarouselColumn{
		ImageURL: "https://example.com/bot/images/item1.jpg",
		Action: &linebot.PostbackAction{
			Label: "A區",
			Data:  "action=click&itemid=0",
		},
	}

	c2 := linebot.ImageCarouselColumn{
		ImageURL: "https://example.com/bot/images/item1.jpg",
		Action: &linebot.PostbackAction{
			Label: "B區",
			Data:  "action=click&itemid=1",
		},
	}
	c_i = append(c_i, &c1, &c2)
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

func Img_Carousel_CampRound_Info(product product.Product) (c_t []*linebot.ImageCarouselColumn) {
	for _, uri := range product.ImageUri {
		fmt.Println("Img_Carousel_CampRound_Info : URI :", uri)
		c1 := linebot.ImageCarouselColumn{
			ImageURL: uri,
			Action: &linebot.PostbackAction{
				Label: product.CampRoundName,
				Data:  "action=click&itemid=0",
			},
		}

		c_t = append(c_t, &c1)
	}
	return c_t
}
