package main

// update logic command
//go:generate goctl api go  -api bitmapServ.api -dir .
// update model command
//go:generate goctl model mysql ddl -c -src ./ddl/bitmapServ.sql -dir ./model
// install swagger plugin
//go:generate  go get -u github.com/zeromicro/goctl-swagger
// update swagger command <需要复制到终端执行>
//go:generate goctl api plugin --plugin goctl-swagger="swagger -filename docs/swagger.json" -api bitMapServ.api -dir .
// see swagger ui
//go:generate docker run --rm -p 8081:8080 -e SWAGGER_JSON=/app/docs/swagger.json -v $PWD:/usr/share/nginx/html/app swaggerapi/swagger-ui

import (
	"flag"
	"fmt"

	"github.com/weblfe/bitmap_service/api/internal/config"
	"github.com/weblfe/bitmap_service/api/internal/handler"
	"github.com/weblfe/bitmap_service/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

const (
	Version = "v1.0.5"
)

var configFile = flag.String("f", "etc/bitMapServ-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("<bitmapServ-%s> server Starting at %s:%d...\n", Version, c.Host, c.Port)
	server.Start()
}
