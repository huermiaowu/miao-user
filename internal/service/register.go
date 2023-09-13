package service

import (
	"context"
	"github.com/huerni/gmitex/pkg/errno"
	"gorm.io/gorm"
	"user/internal/db"

	"user/internal/svc"
	"user/pb"
)

type RegisterService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterService(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterService {
	return &RegisterService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *RegisterService) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// todo: add your logic here and delete this line
	users, err := db.QueryUser(s.ctx, in.Username)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}

	if len(users) > 0 {
		return nil, errno.Err(int(pb.ErrCode_RegisterErrCode), "用户名已存在")
	}
	new_user := db.User{
		Model:     gorm.Model{},
		Username:  in.Username,
		Password:  in.Password,
		Phone:     in.Phone,
		Email:     in.Email,
		AvaUrl:    in.AvaUrl,
		Signature: in.Signature,
	}
	err = db.CreateUser(s.ctx, &new_user)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	userinfo := pb.UserInfo{
		Id:        uint64(new_user.ID),
		Username:  new_user.Username,
		Password:  new_user.Password,
		Email:     new_user.Email,
		Phone:     new_user.Phone,
		AvaUrl:    new_user.AvaUrl,
		Signature: new_user.Signature,
	}
	return &pb.RegisterResp{Userinfo: &userinfo}, nil
}
