package svc

import (
		"github.com/tal-tech/go-zero/core/logx"
		"github.com/tal-tech/go-zero/core/stores/sqlx"
		"github.com/weblfe/bitmap_service/rpc/internal/config"
		"github.com/weblfe/bitmap_service/rpc/internal/model"
)

type ServiceContext struct {
	Config config.Config
	Model model.UserDayLoginModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model: model.NewUserDayLoginModel(sqlx.NewMysql(c.DataSource),c.Table,c.Cache,c),
	}
}

func AssertError(err error) bool  {
	if err==nil {
		return false
	}
	logx.Error(err)
	return true
}
