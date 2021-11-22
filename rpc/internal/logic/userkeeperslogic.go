package logic

import (
		"context"
		"fmt"
		"github.com/tal-tech/go-zero/core/logx"
		"github.com/tal-tech/go-zero/core/stores/sqlx"
		"github.com/weblfe/bitmap_service/rpc/bitMapServ/rpc"
		"github.com/weblfe/bitmap_service/rpc/internal/model"
		"github.com/weblfe/bitmap_service/rpc/internal/svc"
		"github.com/weblfe/bitmap_service/rpc/internal/types"
)

type UserKeepersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserKeepersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserKeepersLogic {
	return &UserKeepersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserKeepersLogic) UserKeepers(request *rpc.RpcMulitRequest) (*rpc.RpcMulitResponse, error) {

	var (
		total int
		left  int
		name  string
		keep  float32
		query model.QueryMap
		data  []*rpc.KeeperStruct
	)
	// 切换数据源
	if request.DataSource != "" {
		l.svcCtx.Model = model.NewUserDayLoginModel(sqlx.NewMysql(l.svcCtx.Config.DataSource), request.DataSource, l.svcCtx.Config.Cache, l.svcCtx.Config)
	}
	for _, day := range request.Days {
		query = model.QueryMap{
			Day:     int(day),
			Channel: request.Channel,
			Role:    request.Role,
			Date:    request.Date,
			Type:    model.ParserDateType(int32(request.Type)),
		}
		name = getName(request.Date, int64(day), request.Channel)
		userArr, err := l.svcCtx.Model.GetUserBitMapArr(query)
		if err != nil {
			data = append(data, &rpc.KeeperStruct{
				Keep:     fmt.Sprintf("%.2f", keep),
				Name:     fmt.Sprintf("%s: %.2f%s", name, keep, "%"),
				Register: 0,
				Left:     0,
			})
			continue
		}
		keep, total, left = GetUserDayKeepByBitmapArr(userArr)
		data = append(data, &rpc.KeeperStruct{
			Keep:     fmt.Sprintf("%.2f", keep),
			Name:     fmt.Sprintf("%s: %.2f%s", name, keep, "%"),
			Register: int64(total),
			Left:     int64(left),
		})
	}
	return &rpc.RpcMulitResponse{
		Code: types.CodeSuccess.GetCode32(),
		Msg:  types.CodeSuccess.GetMsg(),
		Data: data,
	}, nil
}
