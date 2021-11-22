package logic

import (
		"context"
		"fmt"
		"github.com/tal-tech/go-zero/core/stores/sqlx"
		"github.com/weblfe/bitmap_service/rpc/bitMapServ/rpc"
		"github.com/weblfe/bitmap_service/rpc/internal/model"
		"github.com/weblfe/bitmap_service/rpc/internal/svc"
		"github.com/weblfe/bitmap_service/rpc/internal/types"

		"github.com/tal-tech/go-zero/core/logx"
)

type ChannelsUserKeepersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChannelsUserKeepersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelsUserKeepersLogic {
	return &ChannelsUserKeepersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChannelsUserKeepersLogic) ChannelsUserKeepers(req *rpc.RpcMulitChannelsRequest) (*rpc.RpcMulitChannelsResponse, error) {

	var (
		total    = 0
		left     = 0
		name     string
		k        float32
		errArr   []string
		queryMap model.QueryMap
		length   = len(req.Days)
		items    []*rpc.RpcChannelResponse
	)
	// 切换数据源
	if req.DataSource != "" {
		l.svcCtx.Model = model.NewUserDayLoginModel(sqlx.NewMysql(l.svcCtx.Config.DataSource), req.DataSource, l.svcCtx.Config.Cache, l.svcCtx.Config)
	}
	// 多渠道
	for _, channel := range req.Channels {
		it := &rpc.RpcChannelResponse{
			Channel: channel,
			Items:   make([]*rpc.KeeperStruct, length),
		}
		// 多天
		for i, d := range req.Days {
			name = getName(req.Date, int64(d), channel)
			queryMap = model.QueryMap{
				Day:     int(d),
				Channel: channel,
				Role:    req.Role,
				Date:    req.Date,
				Type:    model.ParserDateType(int32(req.Type)),
			}
			userArr, err := l.svcCtx.Model.GetUserBitMapArr(queryMap)
			if err != nil {
				it.Items[i] = &rpc.KeeperStruct{
					Keep:     fmt.Sprintf("%.2f", 0.0),
					Name:     fmt.Sprintf("%s: %.2f%s", name, 0.0, "%"),
					Register: 0,
					Left:     0,
				}
				errArr = append(errArr, err.Error())
			} else {
				k, total, left = GetUserDayKeepByBitmapArr(userArr)
				it.Items[i] = &rpc.KeeperStruct{
					Keep:     fmt.Sprintf("%.2f", k),
					Name:     fmt.Sprintf("%s: %.2f%s", name, k, "%"),
					Register: int64(total),
					Left:     int64(left),
				}
			}
		}
		items = append(items, it)
	}
	if len(errArr) == length && len(errArr) > 0 {
		return &rpc.RpcMulitChannelsResponse{
			Code: types.CodeServerError.GetCode32(),
			Msg:  types.CodeServerError.GetMsg() + " : " + errArr[0],
			Data: items,
		}, nil
	}
	return &rpc.RpcMulitChannelsResponse{
		Code: types.CodeSuccess.GetCode32(),
		Msg:  types.CodeSuccess.GetMsg(),
		Data: items,
	}, nil

}
