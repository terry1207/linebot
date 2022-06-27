package command

type Response16537186462781 struct {
	Success string  `json:"success"`
	Result  Result  `json:"result"`
	Records Records `json:"records"`
}

type Records struct {
	Locations []RecordsLocation `json:"locations"`
}

type RecordsLocation struct {
	DatasetDescription string             `json:"datasetDescription"`
	LocationsName      string             `json:"locationsName"`
	Dataid             string             `json:"dataid"`
	Location           []LocationLocation `json:"location"`
}

type LocationLocation struct {
	LocationName   string           `json:"locationName"`
	Geocode        string           `json:"geocode"`
	Lat            string           `json:"lat"`
	Lon            string           `json:"lon"`
	WeatherElement []WeatherElement `json:"weatherElement"`
}

type WeatherElement struct {
	ElementName string `json:"elementName"`
	Description string `json:"description"`
	Time        []Time `json:"time"`
}

type Time struct {
	StartTime    string         `json:"startTime"`
	EndTime      string         `json:"endTime"`
	ElementValue []ElementValue `json:"elementValue"`
}

type ElementValue struct {
	Value    string   `json:"value"`
	Measures Measures `json:"measures"`
}

type Result struct {
	ResourceID string  `json:"resource_id"`
	Fields     []Field `json:"fields"`
}

type Field struct {
	ID   string `json:"id"`
	Type Type   `json:"type"`
}

type Measures string

const (
	MeasuresNA Measures = "NA "
	Na         Measures = "NA"
	The8方位     Measures = "8方位"
	公尺秒        Measures = "公尺/秒"
	攝氏度        Measures = "攝氏度"
	曝曬級數       Measures = "曝曬級數"
	百分比        Measures = "百分比"
	紫外線指數      Measures = "紫外線指數"
	自定義Ci文字    Measures = "自定義 CI 文字"
	自定義Wx單位    Measures = "自定義 Wx 單位"
	自定義Wx文字    Measures = "自定義 Wx 文字"
	蒲福風級       Measures = "蒲福風級"
)

type Type string

const (
	Double    Type = "Double"
	String    Type = "String"
	Timestamp Type = "Timestamp"
)
