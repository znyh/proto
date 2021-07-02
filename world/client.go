package api

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const AppID = "TODO: ADD APP ID"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (WorldClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return NewWorldClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm --swagger --ecode api.proto

//-conf D:\asrc\zhangyh\goooo\zhub\world\configs -appid service.world -discovery.nodes 172.20.10.2:7171 -deploy.env dev -zone sz001 -grpc.addr 172.20.10.2:9000 -http.perf tcp://0.0.0.0:9020 -log.v 1 -log.dir ./log
