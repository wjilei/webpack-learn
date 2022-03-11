package main

import (
	"encoding/json"
	"net/http"
	"server/src/weather"
	"time"
)

type Result struct {
	V1        []*weather.Dayly
	V2        []*weather.Dayly
	V3        []*weather.Dayly
	V4        []*weather.Dayly
	Aver      []*weather.Dayly
	AverTotal float64
	TotalV1   float64
	TotalV2   float64
	TotalV3   float64
	TotalV4   float64
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/api/get-weather", func(rw http.ResponseWriter, r *http.Request) {
		result := Result{}
		beginTime := time.Date(2018, 12, 1, 0, 0, 0, 0, time.Local)
		year := 1
		for beginTime.Before(time.Date(2022, 12, 1, 0, 0, 0, 0, time.Local)) {
			// log.Println("BeginTime", beginTime)
			startTime := beginTime
			for i := 0; i < 3; i++ {
				// log.Println("startTime", startTime)
				v, err := weather.GetMonthHistoryWeather("101201108", startTime)
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
				l := len(v.Dayly)
				if l == 30 {
					// for _, d := range v.Dayly {
					// 	log.Println(time.Unix(d.Dt, 0))
					// }
					var dd []*weather.Dayly
					dd = append(dd, v.Dayly[:4]...)
					lostPoint := &weather.Dayly{
						Dt: v.Dayly[3].Dt + 86400,
					}
					lostPoint.Temp = v.Dayly[4].Temp
					dd = append(dd, lostPoint)
					dd = append(dd, v.Dayly[4:]...)
					v.Dayly = dd
				}
				// log.Println(len(v.Dayly))
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
		for i, v := range result.V1 {
			dt := &weather.Dayly{
				Dt: v.Dt,
			}
			v1Aver := (v.Temp.Max + v.Temp.Min) / 2
			v2Aver := (result.V2[i].Temp.Max + result.V2[i].Temp.Min) / 2
			v3Aver := (result.V3[i].Temp.Max + result.V3[i].Temp.Min) / 2
			dt.Temp.Max = (v1Aver + v2Aver + v3Aver) / 3
			result.Aver = append(result.Aver, dt)
		}

		var refTemp float32 = 20
		for _, v := range result.V1 {
			averTemp := (v.Temp.Max + v.Temp.Min) / 2
			result.TotalV1 += float64(refTemp - averTemp)
		}
		for _, v := range result.V2 {
			averTemp := (v.Temp.Max + v.Temp.Min) / 2
			result.TotalV2 += float64(refTemp - averTemp)
		}
		for _, v := range result.V3 {
			averTemp := (v.Temp.Max + v.Temp.Min) / 2
			result.TotalV3 += float64(refTemp - averTemp)
		}
		for _, v := range result.V4 {
			averTemp := (v.Temp.Max + v.Temp.Min) / 2
			result.TotalV4 += float64(refTemp - averTemp)
		}
		for _, v := range result.Aver {
			result.AverTotal += float64(refTemp - v.Temp.Max)
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
