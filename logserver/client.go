package api

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"google.golang.org/grpc"
)

const LogServerTopic = "logServerTopic"

// AppID .
const AppID = "TODO: ADD APP ID"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (LogserverClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return NewLogserverClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm --swagger --ecode api.proto
