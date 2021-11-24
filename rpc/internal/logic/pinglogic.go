package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *pb.PingRequest) (*pb.PingResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.PingResponse{}, nil
}
