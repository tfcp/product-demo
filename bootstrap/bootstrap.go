package bootstrap

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"tfpro/app/cron"
	"tfpro/app/process"
	"tfpro/internal/model"
	"tfpro/library/gredis"
	"tfpro/library/log"
	"tfpro/router"
	"tfpro/tools"
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
	//toolArg := gcmd.GetArg(1, "tools")
	opts := gcmd.GetOptAll()
	toolName := gcmd.GetOpt("cmd", "orm")
	switch toolName {
	// ormGen tool
	// -t: tableName
	// -d: database if null is default
	case "orm":
		database := "default"
		table := ""
		if _, ok := opts["d"]; ok {
			database = opts["d"]
		}
		if _, ok := opts["t"]; !ok {
			fmt.Println(`
this a quick orm generate tools.

Usage:

	go run main.go tools -t tableName -d databaseName


				`)
			return
		}
		table = opts["t"]
		tools.OrmGenTools(table, database)
	}

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
