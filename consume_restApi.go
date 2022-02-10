package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Data struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Year         int    `json:"year"`
		Color        string `json:"color"`
		PantoneValue string `json:"pantone_value"`
	} `json:"data"`
	Support struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}

func main2() {
	url := "https://reqres.in/api/products/3"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(response.Body)
	data, err := ioutil.ReadAll(response.Body)
	var responseObject Response
	json.Unmarshal(data, &responseObject)
	log.Println(responseObject)

}
