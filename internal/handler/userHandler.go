// no edit

package handler

import (
	"context"
	"user/internal/service"
	"user/internal/svc"
	"user/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// Ping /api/v1/user/ping
func (s *UserServer) Ping(ctx context.Context, in *pb.PingReq) (*pb.PingResp, error) {
	res, err := service.NewPingService(ctx, s.svcCtx).Ping(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Login /api/v1/user/login
func (s *UserServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	res, err := service.NewLoginService(ctx, s.svcCtx).Login(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Register /api/v1/user/register
func (s *UserServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	res, err := service.NewRegisterService(ctx, s.svcCtx).Register(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetUserInfo /api/v1/user
func (s *UserServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	res, err := service.NewGetUserInfoService(ctx, s.svcCtx).GetUserInfo(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateUserInfo /api/v1/user/update
func (s *UserServer) UpdateUserInfo(ctx context.Context, in *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoResp, error) {
	res, err := service.NewUpdateUserInfoService(ctx, s.svcCtx).UpdateUserInfo(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
