package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rahul10-pu/CIGO0122/handler"
)

func AuthRouter(router *gin.Engine, api handler.Handler) {
	router.POST("/signup", api.SignUp)
	router.POST("/signin", api.SignIn)
}
