package main

import (
	"fmt"

	commonLogic "./commonLogic"
	recruitapi "./recruitapi"
)

type serviceAreaItem struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type largeServiceAreaItem struct {
	APIVersion       string            `json:"api_version"`
	ResultsReturned  string            `json:"results_returned"`
	ResultsStart     int               `json:"results_start"`
	ResultsAvailable int               `json:"results_available"`
	LargeServiceArea []serviceAreaItem `json:"large_service_area"`
}

type resultResponse struct {
	Results largeServiceAreaItem `json:"results"`
}

func main() {
	fmt.Println("hogefuga20181227")
	parameterQuery := make(map[string]string)
	parameterQuery["name"] = "肉"
	parameterQuery["address"] = "東京"

	var queryString = commonLogic.NewSortedQuery(parameterQuery)
	fmt.Println(queryString.String())

	var result = recruitapi.GetBudget("", "json")
	//	var result = recruitapi.GetGourmet("", "json", queryString.String())
	for _, item := range result.Results.Budget {
		fmt.Printf("Name : %s, Code = %s\n", item.Name, item.Code)
	}
}
