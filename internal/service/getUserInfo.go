package service

import (
	"context"
	"fmt"
	"github.com/huerni/gmitex/pkg/errno"
	"user/internal/db"

	"user/internal/svc"
	"user/pb"
)

type GetUserInfoService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoService(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoService {
	return &GetUserInfoService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *GetUserInfoService) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	users, err := db.QueryUserById(s.ctx, in.Id)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	if len(users) == 0 {
		return nil, errno.Err(int(pb.ErrCode_UserNotFoundCode), fmt.Sprintf("无该用户信息"))
	}
	user := users[0]
	userinfo := pb.UserInfo{
		Id:        uint64(user.ID),
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Phone:     user.Phone,
		AvaUrl:    user.AvaUrl,
		Signature: user.Signature,
	}
	return &pb.GetUserInfoResp{Userinfo: &userinfo}, nil
}
