package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
	"time"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	Table      string
	Cache      cache.CacheConf
	SafeTime   time.Duration `json:",default=180s"` // 缓存时长
	DocsPath  string `json:",optional,default=./docs/swagger.json"`
}
