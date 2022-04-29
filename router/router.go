package router

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"tfpro/app/api/demo"
	_ "tfpro/docs" // gin-swagger
	"tfpro/internal/middleware/cors"

	//_ "tfpro/internal/rice"
	//"tfpro/library/utils"
	//rice "github.com/GeertJohan/go.rice"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func RegisterRouter() {
	Router = gin.Default()
	// pprof
	pprof.Register(Router)
	Router.Use(cors.Cors())
	//fs := utils.EmbeddingFileSystem(rice.MustFindBox(enum.RicePath).HTTPBox())
	//Router.Use(utils.Serve("", fs))
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	dm := Router.Group("demo")
	//dm.Use(jwt.JWT()) # jwt auth
	dm.GET("/hello-list", demo.HelloListApi)
	dm.GET("/hello-info", demo.HelloInfoApi)
	dm.GET("/user-list", demo.UserListApi)
	dm.GET("/user-count", demo.UserCountApi)
	dm.GET("/user-detail", demo.UserDetailApi)
	dm.POST("/user-delete", demo.UserDeleteApi)
	dm.POST("/user-change", demo.UserChangeStatusApi)
	dm.POST("/user-save", demo.UserSaveApi)
	us := Router.Group("user")
	us.Any("/login", demo.LoginApi)
	us.GET("/info", demo.InfoApi)
}
