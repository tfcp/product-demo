package client

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"google.golang.org/grpc"
	v1 "tfpro/app/grpc/v1"
	grpcClient "tfpro/library/grpc"
	"tfpro/library/log"
)

var (
	TfRemoteClient v1.RemoteClient
)

func init() {
	var err error
	TfRemoteClient, err = NewTfRemoteClient()
	if err != nil {
		log.Logger.Fatalf("remote-client init error: %v",err)
	}
}

func NewTfRemoteClient() (v1.RemoteClient, error) {
	addr := g.Config().GetString("client.tf_remote_addr")
	opt := grpcClient.DefaultClient()
	gc, err := opt.DialContext(
		context.Background(),
		addr, grpc.WithBlock(),
		// p2c结构
		grpc.WithBalancerName("p2c"),
		// 客户端中间件
		//grpc.WithChainUnaryInterceptor(openmwgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer(), openmwgrpc.LogPayloads())),
	)
	if err != nil {
		return nil, err
	}
	gRpcClient := v1.NewRemoteClient(gc)
	return gRpcClient, nil
}
