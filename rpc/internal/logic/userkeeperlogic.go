package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserKeeperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserKeeperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserKeeperLogic {
	return &UserKeeperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserKeeperLogic) UserKeeper(in *pb.RpcSingleRequest) (*pb.RpcSingleResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.RpcSingleResponse{}, nil
}
