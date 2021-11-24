package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserKeepersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserKeepersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserKeepersLogic {
	return &UserKeepersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserKeepersLogic) UserKeepers(in *pb.RpcMulitRequest) (*pb.RpcMulitResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.RpcMulitResponse{}, nil
}
