package bootstrap

import (
	"fmt"
	"github.com/gogf/gf/os/gcmd"
	"google.golang.org/grpc/reflection"
	"os"
	v1 "tfpro/app/grpc/v1"
	"tfpro/internal/model"
	"tfpro/internal/service/demo"
	"tfpro/library/config"
	"tfpro/library/environment"
	"tfpro/library/gredis"
	"tfpro/library/grpc"
	"tfpro/library/log"
	"tfpro/tools"
)

var (
	// 应用服务引擎，单项目单服务使用唯一Server对象
	gRPCServer = grpc.Default()
)

func bootstrap() {
	var err error
	environment.Setup()
	if err = config.Setup(environment.ConfDriverType); err != nil {
		log.Logger.Fatalf("config init error:%v", err)
		return
	}
	log.Setup()
	if err = gredis.Setup(); err != nil {
		log.Logger.Fatalf("redis init error:%v", err)
	}
	if err = model.Setup(); err != nil {
		log.Logger.Fatalf("db init error:%v", err)
		return
	}
	// 中间件
	//gRPCServer.Use()

	// 应用配置初始化
	if err := gRPCServer.InitServer(); err != nil {
		log.Logger.Fatal(err)
	}
	// 注册grpc server
	v1.RegisterUserServer(gRPCServer.Server, new(demo.RpcServer))

}

func gRpcDebug() {
	// 基于grpcui进行调试
	reflection.Register(gRPCServer.Server)
}

// rpc project
func Run() {
	bootstrap()
	if os.Getenv("ENV") == "DEV" {
		log.Logger.Println("gRpc debug mode open.")
		gRpcDebug()
	}
	address, _ := config.TfConf.Value("grpc.server.addr").String()
	log.Logger.Printf("gRpc server start【 %s 】.", address)
	if err := gRPCServer.Run(address); err != nil {
		log.Logger.Fatal(err)
	}
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

	server             start grpc server.

`
	fmt.Println(helpInfo)

}
