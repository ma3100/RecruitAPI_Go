package main

import (
	"fmt"

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
	fmt.Println("hogefuga201812232")

	var result = recruitapi.GetLargeServiceArea("", "json")
	for _, item := range result.Results.LargeServiceArea {
		fmt.Printf("Name : %s, Code = %s\n", item.Name, item.Code)
	}
}
