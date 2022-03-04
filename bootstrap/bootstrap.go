package bootstrap

import (
	"fmt"
	"gf/cron"
	"gf/internal/model"
	"gf/library/gredis"
	"gf/library/log"
	"gf/process"
	"gf/router"
	"github.com/gogf/gf/frame/g"
)

func bootstrap() {
	log.Setup()
	if err := gredis.Setup(); err != nil {
		log.Logger.Fatalf("redis init error:%v", err)
	}
	if err := model.Setup(); err != nil {
		log.Logger.Fatalf("db init error:%v", err)
		return
	}
}

// web project
func Run() {
	bootstrap()
	router.RegisterRouter()
	addr := g.Config().GetString("api.addr")
	if err := router.Router.Run(addr); err != nil {
		log.Logger.Fatalf("router init error:%v", err)
	}
}

// cronjob
func RunCron() {
	bootstrap()
	cron.Cron()
	select {}
}

// consumer
func RunProcess() {
	bootstrap()
	process.Process()
	select {}
}

func RunTools() {

}

func RunHelp() {
	helpInfo := `
this is a go project. 

Usage:
	
	go run main.go <command>

The commands are:

	server             start http server.
	cron               start the cron job.
	process            start the process.

`
	fmt.Println(helpInfo)

}
