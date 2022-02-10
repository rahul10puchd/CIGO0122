package main

import (
	"encoding/json"
	"fmt"
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
	// var something Something
	something := map[string]interface{}{}
	//json - key and value
	//key is a string
	//value - int, float, bool, string, json - interface

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
func main4() {
	//var user map[int]string //user just got declared, no memory assigned
	user := map[int]interface{}{}
	user[1] = "name"
	user[2] = true
	// fmt.Println(user[2])
	//set
	//does not have duplicate value
	//check or parse in the set
	//use set feature via map
	users := map[string]bool{}
	users["golang"] = true

	// value, ok := users["rahul"]
	// if ok == false {
	// 	fmt.Println("value not there")
	// }

	//fmt.Println(users["rahul"])
	fmt.Println("Hi its my branch")

}
