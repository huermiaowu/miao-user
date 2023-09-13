package service

import (
	"context"

	"user/internal/svc"
	"user/pb"
)

type PingService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingService(ctx context.Context, svcCtx *svc.ServiceContext) *PingService {
	return &PingService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *PingService) Ping(in *pb.PingReq) (*pb.PingResp, error) {
	// todo: add your logic here and delete this line

	return &pb.PingResp{}, nil
}
