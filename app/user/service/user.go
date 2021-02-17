// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package main

import (
	"flag"
	"fmt"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"

	"agent/app/user/service/internal/config"
	"agent/app/user/service/internal/server"
	"agent/app/user/service/internal/svc"
	"agent/app/user/service/user"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(
		c.RpcServerConf, func(grpcServer *grpc.Server) {
			user.RegisterUserServer(grpcServer, srv)
		},
	)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
