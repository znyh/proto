package api

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-kratos/kratos/pkg/log"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"google.golang.org/grpc"
)

const LogServerTopic = "logServerTopic"

// AppID .
const AppID = "logserver-service"
const AppPort = 9000

var (
	_cli LogserverClient
)

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

func DefaultClient() LogserverClient {
	if _cli != nil {
		return _cli
	}

	var err error
	var model = os.Getenv("BOOT_MODE")
	switch model {
	default:
		_cli, err = NewClient(&warden.ClientConfig{
			Dial:    xtime.Duration(time.Second * 3),
			Timeout: xtime.Duration(time.Second * 3),
		})
	}
	if err != nil {
		log.Error("LogserverClient fail, msg:%v", err)
		return nil
	}
	log.Info("init %s client ok .", AppID)
	return _cli
}

func Record(ctx context.Context, req *RecordReq) (resp *RecordRsp) {
	if DefaultClient() == nil {
		return
	}
	resp, err := _cli.Record(ctx, req)
	if err != nil {
		log.Error("Record. err:%v.", err)
		return
	}
	return
}
