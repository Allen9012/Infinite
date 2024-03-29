package main

import (
	"flag"
	"fmt"
	"github.com/Allen9012/Infinite/pkg/interceptors"

	"github.com/Allen9012/Infinite/application/user/rpc/internal/config"
	"github.com/Allen9012/Infinite/application/user/rpc/internal/server"
	"github.com/Allen9012/Infinite/application/user/rpc/internal/svc"
	"github.com/Allen9012/Infinite/application/user/rpc/service"

	"github.com/zeromicro/go-zero/core/conf"
	cs "github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		service.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == cs.DevMode || c.Mode == cs.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	// rpc服务增加拦截器
	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
