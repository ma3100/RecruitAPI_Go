package recruitapi

import (
	"encoding/json"
	"fmt"

	commonLogic "../commonLogic"
)

type resultBudgetResponse struct {
	Results struct {
		ResultsStart     int    `json:"results_start"`
		ResultsReturned  string `json:"results_returned"`
		APIVersion       string `json:"api_version"`
		ResultsAvailable int    `json:"results_available"`
		Budget           []struct {
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"budget"`
	} `json:"results"`
}

func GetBudget(apiKey string, format string) resultBudgetResponse {
	url := "http://webservice.recruit.co.jp/hotpepper/budget/v1/?key=" + apiKey + "&format=" + format

	byteArray := commonLogic.ExecuteGetRequest(url)
	fmt.Println(string(byteArray))

	data := resultBudgetResponse{}
	err := json.Unmarshal(byteArray, &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
