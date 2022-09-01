package line

import (
	"fmt"
	"linebot/internal/model/product"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Search_Time struct {
	Start time.Time
	End   time.Time
}

var Search map[string]*Search_Time

func init() {
	Search = make(map[string]*Search_Time)

	fmt.Println("Init Search ", Search)

	//richmenu.Build_RichMenu()

}

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

				if product, ok := Is_Name_Exist(text_trimspace); ok {
					tmp := Img_Carousel_CampRound_Info(product)
					bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("img carousel",
						&linebot.ImageCarouselTemplate{
							Columns: tmp,
						})).Do()
				}

				switch {

				case text_trimspace == "我要訂營地!":

				case text_trimspace == "營地介紹":
					tmp := Quick_Reply_CampRoundName()
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("選擇分區").WithQuickReplies(&linebot.QuickReplyItems{
						Items: tmp,
					})).Do()

					// default:
					// 	bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text_trimspace)).Do()
					// }
				}
			}
		}

		if event.Type == linebot.EventTypePostback {
			data := Parase_postback(event.Postback.Data)
			switch data.Action {
			case "search":
				value, isExist := Search[event.Source.UserID]
				if !isExist {
					Search[event.Source.UserID] = &Search_Time{}
				}
				switch data.Type {
				case "go":

					bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("date range", linebot.NewButtonsTemplate("", "", "選擇日期",
						linebot.NewDatetimePickerAction("起始日期", "action=search&type=get_start_time", "date", time.Now().Format("2006-01-01"), time.Now().Format("2006-01-01"), ""),
						linebot.NewDatetimePickerAction("結束日期", "action=search&type=get_end_time", "date", time.Now().Format("2006-01-01"), time.Now().Format("2006-01-01"), ""),
					))).Do()
				case "get_start_time":
					str := fmt.Sprintf("起始日期:%s", event.Postback.Params.Date)
					value.Start, _ = time.Parse("2006-01-01", event.Postback.Params.Date)
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(str)).Do()
				case "get_end_time":
					str := fmt.Sprintf("結束日期:%s", event.Postback.Params.Date)
					value.End, _ = time.Parse("2006-01-01", event.Postback.Params.Date)
					fmt.Println(Search[event.Source.UserID])
					delete(Search, event.Source.UserID)
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(str)).Do()

				}

			}
			fmt.Println("data", event.Postback.Data)

		}
	}

}

//快速回覆營位分區名稱
func Quick_Reply_CampRoundName() (q_p []*linebot.QuickReplyButton) {
	fmt.Println("Quick_Reply_CampRoundName")
	products, _ := product.GetAll()
	for _, p := range products {
		tmp := &linebot.QuickReplyButton{
			Action: &linebot.MessageAction{
				Label: p.CampRoundName,
				Text:  p.CampRoundName,
			},
		}
		q_p = append(q_p, tmp)
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

func Add_Carousel_Template() (c_t []*linebot.CarouselColumn) {

	column1 := linebot.CarouselColumn{
		ThumbnailImageURL:    "https://example.com/bot/images/item1.jpg",
		ImageBackgroundColor: "#FFFFFF",
		Title:                "A區",
		Text:                 "5m*5m",
		DefaultAction: &linebot.URIAction{
			Label: "View detail",
			URI:   "https://i.imgur.com/XXwY96T.jpeg",
		},
		Actions: []linebot.TemplateAction{
			&linebot.PostbackAction{
				Label: "訂位",
				Data:  "action=order&itemid=1",
			},
		},
	}
	column2 := linebot.CarouselColumn{
		ThumbnailImageURL:    "https://i.imgur.com/3dthZKo.jpeg",
		ImageBackgroundColor: "#000000",
		Title:                "B區",
		Text:                 "5m*5m",
		DefaultAction: &linebot.URIAction{
			Label: "View detail",
			URI:   "http://example.com/page/222",
		},
		Actions: []linebot.TemplateAction{
			&linebot.PostbackAction{
				Label: "訂位",
				Data:  "action=order&itemid=1",
			},
		},
	}
	c_t = append(c_t, &column1, &column2)

	return c_t
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

func Img_Carousel_CampRound_Info(product product.Product) (c_t []*linebot.ImageCarouselColumn) {
	fmt.Println("Img_Carousel_CampRound_Info", product)
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

type ParseData struct {
	Action string
	Item   int
	Type   string
}

func Parase_postback(data string) (p_d ParseData) {
	str := strings.Split(data, "&")

	for _, p := range str {
		switch {
		case strings.Contains(p, "action"):
			p_d.Action = get_string_data(p)
		case strings.Contains(p, "item"):
			p_d.Item, _ = strconv.Atoi(p)
		case strings.Contains(p, "type"):
			p_d.Type = get_string_data(p)
		}
	}
	return p_d
}

func get_string_data(str string) string {
	i := strings.Index(str, "=")
	tmp := str[i+1:]
	fmt.Println(tmp)
	return tmp
}
