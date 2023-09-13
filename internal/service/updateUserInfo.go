package service

import (
	"context"
	"github.com/huermiaowu/miao-user/internal/db"
	"github.com/huerni/gmitex/pkg/errno"
	"gorm.io/gorm"

	"github.com/huermiaowu/miao-user/internal/svc"
	"github.com/huermiaowu/miao-user/pb"
)

type UpdateUserInfoService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoService(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoService {
	return &UpdateUserInfoService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *UpdateUserInfoService) UpdateUserInfo(in *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoResp, error) {
	// todo: add your logic here and delete this line
	users, err := db.QueryUserById(s.ctx, in.Id)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	if len(users) == 0 {
		return nil, errno.Err(int(pb.ErrCode_UserNotFoundCode), "没有找到该用户")
	}
	user := db.User{
		Model: gorm.Model{
			ID: uint(in.Id),
		},
		Username:  in.Username,
		Password:  in.Password,
		Phone:     in.Phone,
		Email:     in.Email,
		AvaUrl:    in.AvaUrl,
		Signature: in.Signature,
	}

	err = db.UpdateUser(s.ctx, &user)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}

	userinfo := pb.UserInfo{
		Id:        uint64(user.ID),
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Phone:     user.Phone,
		AvaUrl:    user.AvaUrl,
		Signature: user.Signature,
	}
	return &pb.UpdateUserInfoResp{Userinfo: &userinfo}, nil
}
