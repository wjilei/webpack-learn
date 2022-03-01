package weather

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

var (
	ShowapiAppCode = "ccdacb5335ce46c5bae7f811fffdf431"
)

type Dayly struct {
	Dt   int64
	Temp struct {
		Max float32
		Min float32
	}
}

type MonthHistoryWeatherData struct {
	Dayly []*Dayly
}

type getWeatherHistoryResponse struct {
	ShowapiResCode  int    `json:"showapi_res_code,omitempty"`
	ShowapiResError string `json:"showapi_res_error,omitempty"`
	ShowapiResBody  struct {
		List []struct {
			MaxTemperature string `json:"max_temperature,omitempty"`
			MinTemperature string `json:"min_temperature,omitempty"`
			Time           string `json:"time,omitempty"`
		} `json:"list,omitempty"`
	} `json:"showapi_res_body,omitempty"`
}

// GetMonthHistoryWeather 获取d所在月份的历史天气
func GetMonthHistoryWeather(areaId string, d time.Time) (*MonthHistoryWeatherData, error) {
	month := fmt.Sprintf("%d%02d", d.Year(), d.Month())

	s := sling.New().Add("Authorization", "APPCODE "+ShowapiAppCode)
	query := struct {
		Areaid string `url:"areaid,omitempty"`
		Month  string `url:"month,omitempty"`
	}{
		Areaid: areaId,
		Month:  month,
	}
	rsp := getWeatherHistoryResponse{}
	resp, err := s.Get("http://ali-weather.showapi.com/weatherhistory").QueryStruct(&query).ReceiveSuccess(&rsp)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	historyData := MonthHistoryWeatherData{}
	var retDaylyData *Dayly
	for _, d := range rsp.ShowapiResBody.List {
		var year, month, day int
		fmt.Sscanf(d.Time, "%4d%2d%2d", &year, &month, &day)
		var maxTemp, minTemp float32
		fmt.Sscanf(d.MaxTemperature, "%f", &maxTemp)
		fmt.Sscanf(d.MinTemperature, "%f", &minTemp)
		retDaylyData = new(Dayly)
		retDaylyData.Dt = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Unix()
		retDaylyData.Temp.Max = maxTemp
		retDaylyData.Temp.Min = minTemp
		historyData.Dayly = append(historyData.Dayly, retDaylyData)
	}

	return &historyData, nil
}
