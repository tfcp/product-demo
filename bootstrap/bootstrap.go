package bootstrap

import (
	"gf/app/cron"
	"gf/app/process"
	"gf/library/gredis"
	"gf/library/log"
	"gf/router"
	"github.com/gogf/gf/frame/g"
)

// web project
func Run() {
	log.Setup()
	if err := gredis.Setup(); err != nil {
		log.Logger.Fatalf("redis init error:%v", err)
	}
	router.InitRouter()
	addr := g.Config().GetString("api.addr")
	if err := router.Router.Run(addr); err != nil {
		g.Log().Fatal(err)
	}
}

// cronjob
func RunCron() {
	cron.Cron()
	select {}
}

// consumer
func RunProcess() {
	process.Process()
	select {}
}
