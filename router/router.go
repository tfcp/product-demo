package router

import (
	"gf/app/api/demo"
	"gf/app/middleware/cors"
	"gf/app/middleware/jwt"
	//_ "gf/app/statik"
	"gf/library/utils"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func RegisterRouter() {
	Router = gin.Default()
	Router.Use(cors.Cors())
	fs := utils.EmbeddingFileSystem(rice.MustFindBox("../web/dist").HTTPBox())
	Router.Use(utils.Serve("", fs))

	dm := Router.Group("demo")
	dm.Use(jwt.JWT())
	dm.GET("/hello-list", demo.HelloListApi)
	dm.GET("/hello-info", demo.HelloInfoApi)
	dm.Any("/user-list", demo.UserListApi)
	dm.Any("/user-detail", demo.UserDetailApi)
	dm.Any("/user-delete", demo.UserDeleteApi)
	dm.POST("/user-change", demo.UserChangeStatusApi)
	dm.POST("/user-save", demo.UserSaveApi)
	us := Router.Group("user")
	us.Any("/login", demo.LoginApi)
	us.Any("/info", demo.InfoApi)
}
