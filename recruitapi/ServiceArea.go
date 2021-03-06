package recruitapi

import (
	"encoding/json"
	"fmt"

	commonLogic "../commonLogic"
)

type resultServiceAreaResponse struct {
	Results struct {
		APIVersion      string `json:"api_version"`
		ResultsReturned string `json:"results_returned"`
		ResultsStart    int    `json:"results_start"`
		ServiceArea     []struct {
			Name             string `json:"name"`
			LargeServiceArea struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"large_service_area"`
			Code string `json:"code"`
		} `json:"service_area"`
		ResultsAvailable int `json:"results_available"`
	} `json:"results"`
}

func GetServiceArea(apiKey string, format string) resultServiceAreaResponse {
	url := "http://webservice.recruit.co.jp/hotpepper/service_area/v1/?key=" + apiKey + "&format=" + format

	byteArray := commonLogic.ExecuteGetRequest(url)
	fmt.Println(string(byteArray))

	data := resultServiceAreaResponse{}
	err := json.Unmarshal(byteArray, &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
