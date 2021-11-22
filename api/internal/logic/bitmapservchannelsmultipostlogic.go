package logic

import (
		"context"
		"fmt"
		"github.com/tal-tech/go-zero/core/stores/sqlx"
		"github.com/weblfe/bitmap_service/api/model"

		"github.com/weblfe/bitmap_service/api/internal/svc"
		"github.com/weblfe/bitmap_service/api/internal/types"

		"github.com/tal-tech/go-zero/core/logx"
)

type BitMapServChannelsMultiPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBitMapServChannelsMultiPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) BitMapServChannelsMultiPostLogic {
	return BitMapServChannelsMultiPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BitMapServChannelsMultiPostLogic) BitMapServChannelsMultiPost(req types.MulitChannelsRequest) (*types.MulitChannelsResponse, error) {
	// 切换数据源
	if req.DataSource != "" {
		l.svcCtx.Model = model.NewUserDayLoginModel(sqlx.NewMysql(l.svcCtx.Config.DataSource), req.DataSource, l.svcCtx.Config.Cache, l.svcCtx.Config)
	}
	var (
		total  = 0
		left   = 0
		k      float32
		length = len(req.Days)
		items  []*types.ChannelResponse
	)
	// 多渠道
	for _, channel := range req.Channels {
		it := &types.ChannelResponse{
			Channel: channel,
			Items:   make([]*types.KeeperStruct, length),
		}
		// 多天
		for i, d := range req.Days {
			name := getName(req.Date, int64(d), channel)
			userArr, err := l.svcCtx.Model.GetUserBitMapArr(model.QueryMap{Day: d, Channel: channel, Role: req.Role, Date: req.Date, Type: model.ParserDateType(req.Type)})
			if err != nil {
				it.Items[i] = &types.KeeperStruct{
					Keep:  fmt.Sprintf("%.2f", 0.0),
					Name:  fmt.Sprintf("%s: %.2f%s", name, 0.0, "%"),
					Register: 0,
					Left:  0,
				}
			} else {
				k, total, left = GetUserDayKeepByBitmapArr(userArr)
				it.Items[i] = &types.KeeperStruct{
					Keep:  fmt.Sprintf("%.2f", k),
					Name:  fmt.Sprintf("%s: %.2f%s", name, k, "%"),
					Register: total,
					Left:  left,
				}
			}
		}
		items = append(items, it)
	}
	return &types.MulitChannelsResponse{
		BaseResponse: types.BaseResponse{
			Code:    200,
			Msg:     "OK",
			ErrorNo: 0,
		},
		Data: items,
	}, nil
}
