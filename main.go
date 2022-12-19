package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"weather-api/config"
)

type OpenWeatherMapAPIResponse struct {
	Weather []Weather `json:"weather"`
	Dt      int64     `json:"dt"`
}

//OpenWeatherAPIが返す情報の内、時刻と天気の情報を格納する構造体

type Weather struct {
	Main string `json:"main"`
}

//天気の情報の中のMainセクションの情報を格納する構造体

func main() {
	token := config.Config.ApiKey                                 //APIキーを指定(後述)
	city := "Tokyo,jp"                                            //場所を指定
	endPoint := "https://api.openweathermap.org/data/2.5/weather" //APIのエンドポイントを指定(このURLでOK)

	values := url.Values{}     //urlにさらに値を追加する変数を指定
	values.Set("q", city)      //変数cityの値を追加
	values.Set("APPID", token) //変数tokenの値を追加

	res, err := http.Get(endPoint + "?" + values.Encode()) //情報を追加したエンドポイント(URL)にGet
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() //読み込んだBodyを後から閉じる

	bytes, err := ioutil.ReadAll(res.Body) //getで得たBodyを読み込んで変数に代入
	if err != nil {
		panic(err)
	}

	var apiRes OpenWeatherMapAPIResponse // 変数apiResに構造体を代入
	if err := json.Unmarshal(bytes, &apiRes); err != nil {
		panic(err)
	}
	//json形式でBodyの情報を構造体(OpenWeatherMapAPIResponse)に格納し、jsonから構造体の形に整形(json.Unmarshal)

	fmt.Printf("場所:     %v\n", city)
	fmt.Printf("時刻:     %s\n", time.Unix(apiRes.Dt, 0))
	fmt.Printf("天気:     %s\n", apiRes.Weather[0].Main)
	// それぞれ出力
}
