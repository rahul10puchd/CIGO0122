package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Something struct {
	Name    string    `json:"name"`
	Married bool      `json:"married"`
	Age     float64   `json:"age"`
	Address []Address `json:"address"`
	Marks   []int     `json:"marks"`
}
type Address struct {
	Street int    `json:"street"`
	City   string `json:"city"`
}

func main() {
	jsonFile, err := os.Open("something.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	jsonByteValues, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var something Something
	json.Unmarshal(jsonByteValues, &something) //converting json data to struct
	log.Println(something)
	// fmt.Print(string(jsonByteValues)) //converting json data to string
	// newJsonByteValues, err := json.Marshal(something)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// os.WriteFile("some.json",newJsonByteValues)

	//consuming rest api in go code
	//using the rest api
	//returned value ? = json
	//now you must know how to read the json data
	// what is the problem if you read json data in string format
	// read json in strcture format - marshalling , unmarshalling

}
