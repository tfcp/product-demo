package bootstrap

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"syscall"
	"tfpro/app/cron"
	"tfpro/app/process"
	"tfpro/internal/model"
	"tfpro/library/gredis"
	"tfpro/library/log"
	"tfpro/library/zookeeper"
	"tfpro/router"
	"tfpro/tools"
)

func bootstrap() {
	log.Setup()
	if err := gredis.Setup(); err != nil {
		log.Logger.Fatalf("redis init error:%v", err)
	}
	if err := zookeeper.Setup(); err != nil {
		log.Logger.Fatalf("zookeeper init error:%v", err)
		return
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

	// 默认
	//if err := router.Router.Run(addr); err != nil {
	//	log.Logger.Fatalf("router init error:%v", err)
	//}

	// 默认endless服务器会监听下列信号：
	// syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM和syscall.SIGTSTP
	// 接收到 SIGHUP 信号将触发`fork/restart` 实现优雅重启（kill -1 pid会发送SIGHUP信号）
	// 接收到 syscall.SIGINT或syscall.SIGTERM 信号将触发优雅关机
	server := endless.NewServer(addr, router.Router)
	server.BeforeBegin = func(add string) {
		log.Logger.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Logger.Printf("Server err: %v", err)
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
