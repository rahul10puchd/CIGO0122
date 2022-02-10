package main

import (
	"fmt"
	"log"

	"github.com/rahul10-pu/CIGO0122/database"
	"github.com/rahul10-pu/CIGO0122/routers"
)

func init() { //1st
	database.Setup() //establish the database connection
	//database.Example()
	//database.exam()
}
func init() { //2nd
	fmt.Print(1)
}
func init() { //3nd
	fmt.Print(2)
}
func main() {

	engine := routers.Router()          //get the customized engine
	err := engine.Run("127.0.0.1:8080") //start the engine
	if err != nil {
		log.Fatal(err)
	}
}
