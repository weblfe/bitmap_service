// Code generated by goctl. DO NOT EDIT!
// Source: bitMapRpc.proto

package main

//go:generate goctl rpc proto --src bitMapRpc.proto --dir .

import (
	"flag"
	"fmt"

	"github.com/weblfe/bitmap_service/rpc/bitMapServ/rpc"
	"github.com/weblfe/bitmap_service/rpc/internal/config"
	"github.com/weblfe/bitmap_service/rpc/internal/server"
	"github.com/weblfe/bitmap_service/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/bitmapRpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewBitMapServServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		rpc.RegisterBitMapServServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}