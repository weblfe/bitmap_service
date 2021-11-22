package logic

import (
		"context"
		"fmt"
		"github.com/tal-tech/go-zero/core/stores/sqlx"
		"github.com/weblfe/bitmap_service/rpc/internal/model"
		"github.com/weblfe/bitmap_service/rpc/internal/types"

		"github.com/weblfe/bitmap_service/rpc/bitMapServ/rpc"
		"github.com/weblfe/bitmap_service/rpc/internal/svc"

		"github.com/tal-tech/go-zero/core/logx"
)

type UserKeeperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserKeeperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserKeeperLogic {
	return &UserKeeperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserKeeperLogic) UserKeeper(request *rpc.RpcSingleRequest) (*rpc.RpcSingleResponse, error) {
	// 切换数据源
	if request.DataSource != "" {
		l.svcCtx.Model = model.NewUserDayLoginModel(sqlx.NewMysql(l.svcCtx.Config.DataSource), request.DataSource, l.svcCtx.Config.Cache, l.svcCtx.Config)
	}
	var (
		total int
		left  int
		keep  float32
		query = model.QueryMap{
			Day:     int(request.Day),
			Channel: request.Channel,
			Role:    request.Role,
			Date:    request.Date,
			Type:    model.ParserDateType(int32(request.Type)),
		}
		userArr, err = l.svcCtx.Model.GetUserBitMapArr(query)
		name         = getName(request.Date, int64(request.Day), request.Channel)
	)

	if err != nil {
		return &rpc.RpcSingleResponse{
			Code: types.CodeServerError.GetCode32(),
			Msg:  types.CodeServerError.GetMsg() + err.Error(),
		}, nil
	}
	if len(userArr) == 0 {
		return &rpc.RpcSingleResponse{
			Code: types.CodeSuccess.GetCode32(),
			Msg:  types.CodeSuccess.GetMsg(),
			Data: &rpc.KeeperStruct{
				Keep:     fmt.Sprintf("%.2f", keep),
				Name:     fmt.Sprintf("%s: %.2f%s", name, keep, "%"),
				Register: int64(total),
				Left:     int64(left),
			},
		}, nil
	}
	keep, total, left = GetUserDayKeepByBitmapArr(userArr)
	return &rpc.RpcSingleResponse{
		Code: types.CodeSuccess.GetCode32(),
		Msg:  types.CodeSuccess.GetMsg(),
		Data: &rpc.KeeperStruct{
			Keep:     fmt.Sprintf("%.2f", keep),
			Name:     fmt.Sprintf("%s: %.2f%s", name, keep, "%"),
			Register: int64(total),
			Left:     int64(left),
		},
	}, nil
}
