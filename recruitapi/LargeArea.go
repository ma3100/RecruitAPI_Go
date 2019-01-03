package recruitapi

import (
	"encoding/json"
	"fmt"

	commonLogic "../commonLogic"
)

type resultLargeAreaResponse struct {
	Results struct {
		ResultsStart     int    `json:"results_start"`
		ResultsReturned  string `json:"results_returned"`
		APIVersion       string `json:"api_version"`
		ResultsAvailable int    `json:"results_available"`
		LargeArea        []struct {
			ServiceArea struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"service_area"`
			LargeServiceArea struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"large_service_area"`
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"large_area"`
	} `json:"results"`
}

func GetLargeArea(apiKey string, format string) resultLargeAreaResponse {
	url := "http://webservice.recruit.co.jp/hotpepper/large_area/v1/?key=" + apiKey + "&format=" + format

	byteArray := commonLogic.ExecuteGetRequest(url)
	fmt.Println(string(byteArray))

	data := resultLargeAreaResponse{}
	err := json.Unmarshal(byteArray, &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
