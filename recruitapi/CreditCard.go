package recruitapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type resultCreditCardResponse struct {
	Results struct {
		CreditCard []struct {
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"credit_card"`
		ResultsStart     int    `json:"results_start"`
		ResultsReturned  string `json:"results_returned"`
		APIVersion       string `json:"api_version"`
		ResultsAvailable int    `json:"results_available"`
	} `json:"results"`
}

func GetCreditCard(apiKey string, format string) resultCreditCardResponse {
	url := "http://webservice.recruit.co.jp/hotpepper/credit_card/v1/?key=" + apiKey + "&format=" + format

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

	data := resultCreditCardResponse{}
	err = json.Unmarshal(byteArray, &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
