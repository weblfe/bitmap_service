package logic

import (
	"context"

	"github.com/weblfe/bitmap_service/api/internal/svc"
	"github.com/weblfe/bitmap_service/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type BitMapServSingleGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBitMapServSingleGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) BitMapServSingleGetLogic {
	return BitMapServSingleGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BitMapServSingleGetLogic) BitMapServSingleGet(req types.SingleRequest) (*types.SingleResponse, error) {
	var logic = NewBitMapServSingleLogic(l.ctx,l.svcCtx)
	return logic.BitMapServSingle(req)
}
