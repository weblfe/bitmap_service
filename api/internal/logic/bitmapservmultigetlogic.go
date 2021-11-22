package logic

import (
	"context"

	"github.com/weblfe/bitmap_service/api/internal/svc"
	"github.com/weblfe/bitmap_service/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type BitMapServMultiGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBitMapServMultiGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) BitMapServMultiGetLogic {
	return BitMapServMultiGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BitMapServMultiGetLogic) BitMapServMultiGet(req types.MulitRequest) (*types.MulitResponse, error) {
	var logic = NewBitMapServMultiLogic(l.ctx,l.svcCtx)
	return logic.BitMapServMulti(req)
}
