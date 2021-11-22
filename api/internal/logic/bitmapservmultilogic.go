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

type BitMapServMultiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBitMapServMultiLogic(ctx context.Context, svcCtx *svc.ServiceContext) BitMapServMultiLogic {
	return BitMapServMultiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BitMapServMultiLogic) BitMapServMulti(req types.MulitRequest) (*types.MulitResponse, error) {
	if req.DataSource != "" {
		l.svcCtx.Model = model.NewUserDayLoginModel(sqlx.NewMysql(l.svcCtx.Config.DataSource), req.DataSource, l.svcCtx.Config.Cache, l.svcCtx.Config)
	}
	var (
		total = 0
		left  = 0
		k     float32
		items []*types.KeeperStruct
	)

	for _, d := range req.Days {
		name := getName(req.Date, int64(d), req.Channel)
		userArr, err := l.svcCtx.Model.GetUserBitMapArr(model.QueryMap{Day: d, Channel: req.Channel, Role: req.Role, Date: req.Date, Type: model.ParserDateType(req.Type)})
		if err != nil {
			return &types.MulitResponse{
				BaseResponse: types.BaseResponse{
					Code:    200,
					Msg:     err.Error(),
					ErrorNo: 1000,
				},
				Data: nil,
			}, err
		}
		k, total, left = GetUserDayKeepByBitmapArr(userArr)
		items = append(items, &types.KeeperStruct{
			Keep:  fmt.Sprintf("%.2f", k),
			Name:  fmt.Sprintf("%s: %.2f%s", name, k, "%"),
			Register: total,
			Left:  left,
		})
	}

	return &types.MulitResponse{
		BaseResponse: types.BaseResponse{
			Code:    200,
			Msg:     "OK",
			ErrorNo: 0,
		},
		Data: items,
	}, nil
}
