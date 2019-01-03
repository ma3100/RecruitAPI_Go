package recruitapi

import (
	"encoding/json"
	"fmt"

	commonLogic "../commonLogic"
)

type resultGenreResponse struct {
	Results struct {
		ResultsStart     int    `json:"results_start"`
		ResultsReturned  string `json:"results_returned"`
		APIVersion       string `json:"api_version"`
		ResultsAvailable int    `json:"results_available"`
		Genre            []struct {
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"genre"`
	} `json:"results"`
}

func GetGenre(apiKey string, format string) resultGenreResponse {
	url := "http://webservice.recruit.co.jp/hotpepper/genre/v1/?key=" + apiKey + "&format=" + format

	byteArray := commonLogic.ExecuteGetRequest(url)
	fmt.Println(string(byteArray))

	data := resultGenreResponse{}
	err := json.Unmarshal(byteArray, &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
