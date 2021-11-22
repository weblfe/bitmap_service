package logic

import (
	"context"
	"time"

	"github.com/weblfe/bitmap_service/rpc/bitMapServ/rpc"
	"github.com/weblfe/bitmap_service/rpc/internal/svc"

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

func (l *PingLogic) Ping(request *rpc.PingRequest) (*rpc.PingResponse, error) {
	return &rpc.PingResponse{
		Timestamp: time.Now().Unix(),
		Ack:       request.Ack + 1,
	}, nil
}
