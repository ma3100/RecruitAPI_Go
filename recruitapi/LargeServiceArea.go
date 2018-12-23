package recruitapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func AppendString(n int) string {
	base := "abc"
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, base...)
	}
	return string(buf)
}

// 先頭大文字だと外部から参照できる
func GetLargeServiceArea(apiKey string, format string) resultResponse {
	url := "http://webservice.recruit.co.jp/hotpepper/large_service_area/v1/?key=" + apiKey + "&format=" + format

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

	data := resultResponse{}
	err = json.Unmarshal(byteArray, &data)
	if err != nil {
		fmt.Println(err)
	}

	return data
}
