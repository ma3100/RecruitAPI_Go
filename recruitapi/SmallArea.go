package recruitapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type resultSmallAreaResponse struct {
	Results struct {
		SmallArea []struct {
			ServiceArea struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"service_area"`
			MiddleArea struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"middle_area"`
			LargeServiceArea struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"large_service_area"`
			LargeArea struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"large_area"`
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"small_area"`
		ResultsStart     int    `json:"results_start"`
		ResultsReturned  string `json:"results_returned"`
		APIVersion       string `json:"api_version"`
		ResultsAvailable int    `json:"results_available"`
	} `json:"results"`
}

func GetSmallArea(apiKey string, format string) resultSmallAreaResponse {
	url := "http://webservice.recruit.co.jp/hotpepper/service_area/v1/?key=" + apiKey + "&format=" + format

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	byteArray, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(byteArray))

	data := resultSmallAreaResponse{}
	err = json.Unmarshal(byteArray, &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
