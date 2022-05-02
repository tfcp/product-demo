package demo

import (
	"context"
	v1 "tfpro/app/grpc/v1"
)

func (this *RpcServer) Info(context.Context, *v1.InfoReq) (*v1.InfoRes, error) {
	res := v1.InfoRes{
		Name: "test",
		Age:  24,
	}
	return &res, nil
}

func (this *RpcServer) List(context.Context, *v1.ListReq) (*v1.ListRes, error) {
	return &v1.ListRes{}, nil
}
