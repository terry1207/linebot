package call_weath_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	"linebot/internal/application/command"
)

type Weath interface {
	GetWeather() command.Response16537186462781
}

type Weath_Api_Input struct {
	Route         string `json:"Route"`
	DataId        string `json:"dataid"`
	Authorization string `json:"Authorization"`
	Limit         int    `json:"limit"`
	Offset        int    `json:"offset"`
	Format        string `json:"format" `
	LocationName  string `json:"locationName"`
	ElementName   string `json:"elementName"`
	Sort          string `json:"sort"`
	StartTime     string `json:"startTime"`
	TimeFrom      string `json:"timeFrom"`
	TimeTo        string `json:"timeTo"`
}

var (
	token string = "CWB-096B3BC7-414C-4C11-A879-DBF28C469749"
	route string = "https://opendata.cwb.gov.tw/api/v1/rest/datastore/"
)

//combine full api route
func FullRoute(c Weath_Api_Input) string {
	var fullroute string
	c.Authorization = token
	c.Route = route

	t := reflect.TypeOf(c)
	v := reflect.ValueOf(c)

	joinstring := func(query, value string) {
		switch query {
		case "Authorization":
			fullroute = fmt.Sprintf("%s?%s=%s", fullroute, query, value)
		case "Route":
			fullroute = value
		case "dataid":
			fullroute = fmt.Sprintf("%s%s", fullroute, value)
		default:
			fullroute = fmt.Sprintf("%s&%s=%s", fullroute, query, value)

		}

	}

	for i := 0; i < t.NumField(); i++ {
		switch v.Field(i).Kind().String() {

		case "int":
			if v.Field(i).Int() != 0 {
				joinstring(t.Field(i).Tag.Get("json"), strconv.Itoa(int(v.Field(i).Int())))

			}
		case "string":
			if v.Field(i).String() != "" {
				joinstring(t.Field(i).Tag.Get("json"), v.Field(i).String())

			}
		}

		// if v.Field(i).Elem().IsNil() {
		// 	fmt.Println(t.Field(i))
		// }

	}

	return fullroute

}

//查詢三天 每12小時降雨機率
func Search_Hour_12_PoP(c command.Response16537186462781) string {
	var v string
	m := c.Records.Locations[0]

	v = fmt.Sprintf("%s %s %s\n", m.LocationsName, m.Location[0].LocationName, m.Location[0].WeatherElement[0].Description)

	for _, m := range m.Location[0].WeatherElement[0].Time {
		v = fmt.Sprintf("%s %s ~ %s 降雨機率： %s %s \n", v, m.StartTime, m.EndTime, m.ElementValue[0].Value, "%")
	}
	fmt.Println(v)
	return v
}

//呼叫 氣象局 api
func (c Weath_Api_Input) GetWeather() command.Response16537186462781 {
	fullroute := FullRoute(c)
	fmt.Println(fullroute)
	resq, err := http.Get(fullroute)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resq.Body)
	//fmt.Println(string(body))

	var mapresult command.Response16537186462781
	err1 := json.Unmarshal(body, &mapresult)
	if err1 != nil {
		fmt.Println("jsontomap fail")
	}

	return mapresult
}

//對照縣市 回傳 {dataid}
func Return_location_code_2days(location string) string {
	var dataid string
	switch location {
	case "宜蘭縣":
		dataid = "F-D0047-001"
	case "桃園市":
		dataid = "F-D0047-005"
	case "新竹縣":
		dataid = "F-D0047-009"
	case "苗栗縣":
		dataid = "F-D0047-013"
	case "彰化縣":
		dataid = "F-D0047-017"
	case "南投縣":
		dataid = "F-D0047-021"
	case "雲林縣":
		dataid = "F-D0047-025"
	case "嘉義縣":
		dataid = "F-D0047-029"
	case "屏東縣":
		dataid = "F-D0047-033"
	case "臺東縣":
		dataid = "F-D0047-037"
	case "花蓮縣":
		dataid = "F-D0047-041"
	case "澎湖縣":
		dataid = "F-D0047-045"
	case "基隆市":
		dataid = "F-D0047-049"
	case "新竹市":
		dataid = "F-D0047-053"
	case "嘉義市":
		dataid = "F-D0047-057"
	case "臺北市":
		dataid = "F-D0047-061"
	case "高雄市":
		dataid = "F-D0047-065"
	case "新北市":
		dataid = "F-D0047-069"
	case "臺中市":
		dataid = "F-D0047-073"
	case "臺南市":
		dataid = "F-D0047-077"
	case "連江縣":
		dataid = "F-D0047-081"
	case "金門縣":
		dataid = "F-D0047-085"
	}
	fmt.Println(dataid)
	return dataid
}

func Get(c Weath) command.Response16537186462781 {

	return c.GetWeather()
}
