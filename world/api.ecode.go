// Code generated by protoc-gen-ecode v0.1, DO NOT EDIT.
// source: api.proto

/*
Package api is a generated ecode package.
This code was generated with kratos/tool/protobuf/protoc-gen-ecode v0.1.

package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

It is generated from these files:
	api.proto
*/
package api

import (
	"github.com/go-kratos/kratos/pkg/ecode"
)

// to suppressed 'imported but not used warning'
var _ ecode.Codes

// WorldErrCode ecode
var (
	RoomNumNotExist           = ecode.New(100200)
	ParseRoomConfigJsonFailed = ecode.New(100201)
	RoomCardsCountError       = ecode.New(100202)
	GameIdNotConfigured       = ecode.New(100203)
	RoomCardsUseFailed        = ecode.New(100204)
	RocoverRoomCardsFailed    = ecode.New(100205)
	RoomCardsConfigNotFound   = ecode.New(100206)
	SetRoomConfigFail         = ecode.New(100207)
	CreateRoomNumFailed       = ecode.New(100208)
	NoFreeRoomNumsFail        = ecode.New(100209)
	ApplyCustomRoomNumFail    = ecode.New(100210)
	DisbandCustomRoomNumFail  = ecode.New(100211)
)