package logic

import (
		"context"
		"fmt"
		"github.com/tal-tech/go-zero/core/stores/sqlx"
		"github.com/weblfe/bitmap_service/api/internal/svc"
		"github.com/weblfe/bitmap_service/api/internal/types"
		"github.com/weblfe/bitmap_service/api/model"

		"github.com/tal-tech/go-zero/core/logx"
)

type BitMapServSingleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBitMapServSingleLogic(ctx context.Context, svcCtx *svc.ServiceContext) BitMapServSingleLogic {
	return BitMapServSingleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BitMapServSingleLogic) BitMapServSingle(req types.SingleRequest) (*types.SingleResponse, error) {
	if req.DataSource != "" {
		l.svcCtx.Model = model.NewUserDayLoginModel(sqlx.NewMysql(l.svcCtx.Config.DataSource), req.DataSource, l.svcCtx.Config.Cache, l.svcCtx.Config)
	}
	var (
		total         = 0
		left          = 0
		keep  float32 = 0.00
		name          = getName(req.Date, int64(req.Day), req.Channel)
	)
	userArr, err := l.svcCtx.Model.GetUserBitMapArr(model.QueryMap{Day: req.Day, Channel: req.Channel, Role: req.Role, Date: req.Date, Type: model.ParserDateType(req.Type)})
	if err != nil {
		return &types.SingleResponse{
			BaseResponse: types.BaseResponse{
				Code:    200,
				Msg:     err.Error(),
				ErrorNo: 500,
			},
			Data: nil,
		}, nil
	}

	if len(userArr) == 0 {
		return &types.SingleResponse{
			BaseResponse: types.BaseResponse{
				Code: 200,
				Msg:  "OK",
			},
			Data: &types.KeeperStruct{
				Keep:  fmt.Sprintf("%.2f", keep),
				Name:  fmt.Sprintf("%s: %.2f%s", name, keep, "%"),
				Register: total,
				Left:  left,
			},
		}, nil
	}

	keep, total, left = GetUserDayKeepByBitmapArr(userArr)
	return &types.SingleResponse{
		BaseResponse: types.BaseResponse{
			Code: 200,
			Msg:  "OK",
		},
		Data: &types.KeeperStruct{
			Keep:  fmt.Sprintf("%.2f", keep),
			Name:  fmt.Sprintf("%s: %.2f%s", name, keep, "%"),
			Register: total,
			Left:  left,
		},
	}, nil
}
