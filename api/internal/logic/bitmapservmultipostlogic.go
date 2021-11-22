package logic

import (
	"context"

	"github.com/weblfe/bitmap_service/api/internal/svc"
	"github.com/weblfe/bitmap_service/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type BitMapServMultiPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBitMapServMultiPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) BitMapServMultiPostLogic {
	return BitMapServMultiPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BitMapServMultiPostLogic) BitMapServMultiPost(req types.MulitRequest) (*types.MulitResponse, error) {
	var logic = NewBitMapServMultiLogic(l.ctx,l.svcCtx)
	return logic.BitMapServMulti(req)
}
