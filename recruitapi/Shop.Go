package recruitapi

import (
	"encoding/json"
	"fmt"

	commonLogic "../commonLogic"
)

type resultShopResponse struct {
	Results struct {
		ResultsStart    int    `json:"results_start"`
		ResultsReturned string `json:"results_returned"`
		APIVersion      string `json:"api_version"`
		Shop            []struct {
			NameKana string `json:"name_kana"`
			Name     string `json:"name"`
			Urls     struct {
				Pc string `json:"pc"`
			} `json:"urls"`
			Genre struct {
				Name string `json:"name"`
			} `json:"genre"`
			Desc    string `json:"desc"`
			Address string `json:"address"`
			ID      string `json:"id"`
		} `json:"shop"`
		ResultsAvailable int `json:"results_available"`
	} `json:"results"`
}

func GetShop(apiKey string, format string, parameterString string) resultShopResponse {
	url := "http://webservice.recruit.co.jp/hotpepper/shop/v1/?key=" + apiKey + "&format=" + format + "&" + parameterString

	byteArray := commonLogic.ExecuteGetRequest(url)
	fmt.Println(string(byteArray))

	data := resultShopResponse{}
	err := json.Unmarshal(byteArray, &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
