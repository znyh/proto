package api

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const AppID = "share-service"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (ShareClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return NewShareClient(cc), nil
}

// -conf D:\asrc\zhangyh\middle-end\share\configs

// ็ๆ gRPC ไปฃ็ 
//go:generate kratos tool protoc --grpc --bm --swagger --ecode api.proto
