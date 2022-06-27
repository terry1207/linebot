package user_search

import (
	"fmt"

	"linebot/internal/application/command"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Search_Weath struct {
	Search_Judge
	Search_Record command.Location
	Search_Input  User_Search_Weath
}

type Search_Judge struct {
	Start_Search      bool
	Search_Region     bool
	Search_City       bool
	Search_Town_Index int
}

type Page string

const (
	Next_Page     Page = "下一頁"
	Previous_Page Page = "上一頁"
)

func Quick_Reply_Region() linebot.QuickReplyItems {
	var c linebot.QuickReplyItems

	for i := 0; i < len(command.Location_List.Regions); i++ {
		c.Items = append(c.Items, &linebot.QuickReplyButton{
			Action: &linebot.MessageAction{
				Label: command.Location_List.Regions[i].RigionName,
				Text:  command.Location_List.Regions[i].RigionName,
			},
		})
	}

	fmt.Println(c.Items)
	return c
}
func (t *Search_Weath) Quick_Reply_City() linebot.QuickReplyItems {
	var c linebot.QuickReplyItems

	for k, m := range command.Location_List.Regions {
		if m.RigionName == t.Search_Input.RegionName {

			p := command.Location_List.Regions[k]
			t.Search_Record.Regions = append(t.Search_Record.Regions, p)
			for i := 0; i < len(p.Citys); i++ {
				c.Items = append(c.Items, &linebot.QuickReplyButton{
					Action: &linebot.MessageAction{
						Label: p.Citys[i].CityName,
						Text:  p.Citys[i].CityName,
					},
				})
			}
		}
	}

	return c
}

func (t *Search_Weath) Quick_Reply_Town() linebot.QuickReplyItems {
	var c linebot.QuickReplyItems
	for k, m := range t.Search_Record.Regions[0].Citys {
		if t.Search_Input.Location == m.CityName {

			t.Search_Record.Regions[0].Citys = []command.City{t.Search_Record.Regions[0].Citys[k]}

			var (
				index_lower = t.Search_Town_Index * 10
				index_upper = index_lower + 10
				next        bool
				previous    bool
			)

			switch {
			case index_lower == 0:
				if len(m.Towns) < index_upper {
					index_upper = len(m.Towns)

				} else {
					next = true
				}
			case index_lower > 0:
				previous = true
				if len(m.Towns) < index_upper {
					index_upper = len(m.Towns)
					next = false
				} else {
					next = true
				}

			}

			if previous {
				c.Items = append(c.Items, &linebot.QuickReplyButton{
					Action: &linebot.MessageAction{
						Label: string(Previous_Page),
						Text:  string(Previous_Page),
					},
				})
			}

			for i := index_lower; i < index_upper; i++ {
				c.Items = append(c.Items, &linebot.QuickReplyButton{
					Action: &linebot.MessageAction{
						Label: m.Towns[i],
						Text:  m.Towns[i],
					},
				})
			}

			if next {
				c.Items = append(c.Items, &linebot.QuickReplyButton{
					Action: &linebot.MessageAction{
						Label: string(Next_Page),
						Text:  string(Next_Page),
					},
				})
			}
			fmt.Println(index_lower, index_upper, next, previous)

		}

	}
	fmt.Println("range", t.Search_Record.Regions)

	fmt.Println(t)
	return c
}
