package test

import (
	"github.com/gogf/gf/frame/g"
	v1 "tfpro/app/grpc/v1"
	"tfpro/internal/model"
	"tfpro/internal/service/demo"
	"tfpro/library/gredis"
	"tfpro/library/grpc"
	"tfpro/library/log"
)

var (
	TestgRPCServer = grpc.Default()
)

func SetupServer() {
	g.Config().SetPath("../../../config")
	log.Setup()
	if err := gredis.Setup(); err != nil {
		log.Logger.Fatalf("redis init error:%v", err)
	}
	if err := model.Setup(); err != nil {
		log.Logger.Fatalf("db init error:%v", err)
		return
	}
	// 应用配置初始化
	if err := TestgRPCServer.InitServer(); err != nil {
		log.Logger.Fatal(err)
	}
	// 注册grpc server
	v1.RegisterUserServer(TestgRPCServer.Server, new(demo.RpcServer))
}
