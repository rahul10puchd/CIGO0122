package main

import (
	"log"

	"github.com/rahul10-pu/CIGO0122/database"
	"github.com/rahul10-pu/CIGO0122/routers"
)

func main() {
	database.Setup()                    //establish the database connection
	engine := routers.Router()          //get the customized engine
	err := engine.Run("127.0.0.1:8080") //start the engine
	if err != nil {
		log.Fatal(err)
	}
}
