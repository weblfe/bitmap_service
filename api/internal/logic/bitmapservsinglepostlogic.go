package logic

import (
	"context"

	"github.com/weblfe/bitmap_service/api/internal/svc"
	"github.com/weblfe/bitmap_service/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type BitMapServSinglePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBitMapServSinglePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) BitMapServSinglePostLogic {
	return BitMapServSinglePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BitMapServSinglePostLogic) BitMapServSinglePost(req types.SingleRequest) (*types.SingleResponse, error) {
	var logic = NewBitMapServSingleLogic(l.ctx,l.svcCtx)
	return logic.BitMapServSingle(req)
}
