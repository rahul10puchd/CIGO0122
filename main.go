package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rahul10-pu/CIGO0122/database"
	"github.com/rahul10-pu/CIGO0122/handler"
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
func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func AuthMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	// Foo()
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL)
		token := c.GetHeader("token")
		fmt.Println("got token:	" + token)
		isValid, err := handler.ValidateToken(token)
		if err != nil && !isValid {
			respondWithError(c, 401, "Invalid API token")
			return
		}
		c.Next()
	}
}
func main() {
	engine := gin.Default() //get the default engine for further customizatikon
	api := handler.Handler{
		DB: database.GetDB(), //set the handler DB
	} //get the customized engine
	engine.Use(AuthMiddleware())
	routers.BookRouter(engine, api)
	routers.AuthRouter(engine, api)

	err := engine.Run("127.0.0.1:8080") //start the engine
	if err != nil {
		log.Fatal(err)
	}
}
