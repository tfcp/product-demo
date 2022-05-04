package demo

import (
	"context"
	v1 "tfpro/app/grpc/v1"
)

func (this *RpcServer) Info(ctx context.Context, req *v1.InfoReq) (*v1.InfoRes, error) {
	res := v1.InfoRes{
		Name:   "test-info",
		Age:    1,
		UserId: req.UserId,
	}
	//client.TfRemoteClient.List(context.Background(),)
	return &res, nil
}

func (this *RpcServer) List(ctx context.Context, req *v1.ListReq) (*v1.ListRes, error) {
	resList := new(v1.ListRes)
	resList.List = make([]*v1.InfoRes, 0)
	info := v1.InfoRes{
		Name:   "test-list",
		Age:    2,
		UserId: req.UserId,
	}
	resList.List = append(resList.List,&info)
	return resList, nil
}
