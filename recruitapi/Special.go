package recruitapi

import (
	"encoding/json"
	"fmt"

	commonLogic "../commonLogic"
)

type resultSpecialResponse struct {
	Results struct {
		ResultsStart    int    `json:"results_start"`
		ResultsReturned string `json:"results_returned"`
		APIVersion      string `json:"api_version"`
		Special         []struct {
			SpecialCategory struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"special_category"`
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"special"`
		ResultsAvailable int `json:"results_available"`
	} `json:"results"`
}

func GetSpecial(apiKey string, format string, parameterString string) resultSpecialResponse {
	url := "http://webservice.recruit.co.jp/hotpepper/special/v1/?key=" + apiKey + "&format=" + format

	byteArray := commonLogic.ExecuteGetRequest(url)
	fmt.Println(string(byteArray))

	data := resultSpecialResponse{}
	err := json.Unmarshal(byteArray, &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
