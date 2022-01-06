package router

import (
	"gf/app/api/demo"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()
	dm := Router.Group("demo")
	dm.GET("/hello", demo.Hello)
}
