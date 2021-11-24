package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChannelsUserKeepersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChannelsUserKeepersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelsUserKeepersLogic {
	return &ChannelsUserKeepersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChannelsUserKeepersLogic) ChannelsUserKeepers(in *pb.RpcMulitChannelsRequest) (*pb.RpcMulitChannelsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.RpcMulitChannelsResponse{}, nil
}
