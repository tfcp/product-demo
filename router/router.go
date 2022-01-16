package router

import (
	"gf/app/api/demo"
	"gf/app/api/user"
	"gf/app/middleware/cors"
	"gf/app/middleware/jwt"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()
	Router.Use(cors.Cors())

	dm := Router.Group("demo")
	dm.Use(jwt.JWT())
	dm.GET("/hello-list", demo.HelloListApi)
	dm.GET("/hello-info", demo.HelloInfoApi)
	dm.Any("/user-list", demo.UserListApi)
	us := Router.Group("user")
	us.Any("/login", user.LoginApi)
	us.Any("/info", user.InfoApi)
}
