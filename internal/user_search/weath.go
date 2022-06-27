package user_search

import (
	"fmt"

	"linebot/internal/call_weath_api"
)

type User_Search_Weath struct {
	RegionName   string
	Location     string
	LocationName string
	ElementName  string
	TimeFrom     string
	TimeTo       string
}

type Hour_12_PoP struct {
	City        string
	Town        string
	Description string
	Data        []string
}

var ID_Search map[string]*Search_Weath

func init() {
	tmp := &ID_Search
	*tmp = make(map[string]*Search_Weath)

	fmt.Println("Init", ID_Search)

	//richmenu.Build_RichMenu()

}

//處理user 輸入資訊 執行 call api 回傳每12小時降雨機率
func Call_Weath_Api(c *User_Search_Weath) string {
	var input call_weath_api.Weath_Api_Input
	input.DataId = call_weath_api.Return_location_code_2days(c.Location)
	input.LocationName = c.LocationName
	input.ElementName = c.ElementName
	input.Sort = "time"
	input.TimeFrom = c.TimeFrom
	input.TimeTo = c.TimeTo
	value := call_weath_api.Get(input)

	return call_weath_api.Search_Hour_12_PoP(value)
}
