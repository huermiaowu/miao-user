package router

import (
	"context"
	"github.com/huerni/gmitex/pkg/gw"
)

type UserRouter struct {
	Paths []string
}

func NewUserRouter() *UserRouter {
	paths := make([]string, 0)

	paths = append(paths, "/api/v1/user")

	return &UserRouter{Paths: paths}
}

func (ur *UserRouter) AddRouters(ctx context.Context, client gw.GatewayClient, info any) error {
	return client.AddRoute(ctx, info)
}
