package logic

import (
	"context"

	"github.com/weblfe/bitmap_service/api/internal/svc"
	"github.com/weblfe/bitmap_service/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type BitMapServChannelsMultiGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBitMapServChannelsMultiGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) BitMapServChannelsMultiGetLogic {
	return BitMapServChannelsMultiGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BitMapServChannelsMultiGetLogic) BitMapServChannelsMultiGet(req types.MulitChannelsRequest) (*types.MulitChannelsResponse, error) {
	var logic = NewBitMapServChannelsMultiPostLogic(l.ctx, l.svcCtx)
	return logic.BitMapServChannelsMultiPost(req)
}
