package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rahul10-pu/CIGO0122/database"
	"github.com/rahul10-pu/CIGO0122/handler"
)

func Router() *gin.Engine {
	router := gin.Default() //get the default engine for further customizatikon
	api := handler.Handler{
		DB: database.GetDB(), //set the handler DB
	}

	router.GET("/books", api.GetBooks) //set the function for http://localhost:8080/books : Get request
	//while calling handler function, gin will pass *gin.Context in the handler function
	return router
}
