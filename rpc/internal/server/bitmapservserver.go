// Code generated by goctl. DO NOT EDIT!
// Source: bitMapRpc.proto

package server

import (
	"context"

	"rpc/internal/logic"
	"rpc/internal/svc"
	"rpc/pb"
)

type BitMapServServer struct {
	svcCtx *svc.ServiceContext
}

func NewBitMapServServer(svcCtx *svc.ServiceContext) *BitMapServServer {
	return &BitMapServServer{
		svcCtx: svcCtx,
	}
}

func (s *BitMapServServer) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *BitMapServServer) UserKeeper(ctx context.Context, in *pb.RpcSingleRequest) (*pb.RpcSingleResponse, error) {
	l := logic.NewUserKeeperLogic(ctx, s.svcCtx)
	return l.UserKeeper(in)
}

func (s *BitMapServServer) UserKeepers(ctx context.Context, in *pb.RpcMulitRequest) (*pb.RpcMulitResponse, error) {
	l := logic.NewUserKeepersLogic(ctx, s.svcCtx)
	return l.UserKeepers(in)
}

func (s *BitMapServServer) ChannelsUserKeepers(ctx context.Context, in *pb.RpcMulitChannelsRequest) (*pb.RpcMulitChannelsResponse, error) {
	l := logic.NewChannelsUserKeepersLogic(ctx, s.svcCtx)
	return l.ChannelsUserKeepers(in)
}
