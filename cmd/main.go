package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/huermiaowu/miao-user/internal/app"
	"github.com/huermiaowu/miao-user/internal/config"
	"github.com/huermiaowu/miao-user/internal/handler"
	"github.com/huermiaowu/miao-user/internal/svc"
	user "github.com/huermiaowu/miao-user/pb"
	"github.com/huerni/gmitex/pkg/http/handlers"
	"github.com/huerni/gmitex/pkg/http/response"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	c, err := config.InitConfig("etc/cfg.toml")
	if err != nil {
		panic(err)
	}

	g := app.NewGmServer(c, user.RegisterUserHandlerFromEndpoint, func(server *grpc.Server) {
		ctx := svc.NewServiceContext(c)
		srv := handler.NewUserServer(ctx)
		user.RegisterUserServer(server, srv)
	})

	g.HttpServer.AddMuxOp(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &response.CustomMarshaler{
				M: &runtime.JSONPb{
					MarshalOptions:   protojson.MarshalOptions{},
					UnmarshalOptions: protojson.UnmarshalOptions{},
				}}}),
		runtime.WithErrorHandler(handlers.ErrorHandler),
	)

	g.Start(context.Background())
	g.WaitForShutdown(context.Background())
}
