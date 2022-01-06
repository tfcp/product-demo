package bootstrap

import (
	"gf/router"
	"github.com/gogf/gf/frame/g"
)

func Run() {
	router.InitRouter()
	addr := g.Config().GetString("api.addr")
	if err := router.Router.Run(addr); err != nil {
		g.Log().Fatal(err)
	}
}
