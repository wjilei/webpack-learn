package main

import (
	"encoding/json"
	"net/http"
	"server/src/weather"
	"time"
)

type Result struct {
	V1 []*weather.Dayly
	V2 []*weather.Dayly
	V3 []*weather.Dayly
	V4 []*weather.Dayly
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/api/get-weather", func(rw http.ResponseWriter, r *http.Request) {
		result := Result{}
		beginTime := time.Date(2018, 12, 1, 0, 0, 0, 0, time.Local)
		year := 1
		for beginTime.Before(time.Date(2022, 12, 1, 0, 0, 0, 0, time.Local)) {
			startTime := beginTime
			for i := 0; i < 3; i++ {
				v, err := weather.GetMonthHistoryWeather("101210114", startTime)
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
				if year == 1 {
					result.V1 = append(result.V1, v.Dayly...)
				} else if year == 2 {
					result.V2 = append(result.V2, v.Dayly...)
				} else if year == 3 {
					result.V3 = append(result.V3, v.Dayly...)
				} else if year == 4 {
					result.V4 = append(result.V4, v.Dayly...)
				}

				startTime = startTime.AddDate(0, 1, 0)
			}
			beginTime = beginTime.AddDate(1, 0, 0)
			year++
		}

		b, err := json.Marshal(result)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Write(b)
	})
	http.ListenAndServe(":8099", mux)
}
