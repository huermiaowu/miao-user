package service

import (
	"context"
	"github.com/huermiaowu/miao-user/internal/db"
	"github.com/huermiaowu/miao-user/internal/svc"
	"github.com/huermiaowu/miao-user/pb"
	"github.com/huerni/gmitex/pkg/errno"
)

type LoginService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginService(ctx context.Context, svcCtx *svc.ServiceContext) *LoginService {
	return &LoginService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *LoginService) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line
	users, err := db.QueryUser(s.ctx, in.Username)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	if len(users) == 0 {
		return nil, errno.Err(int(pb.ErrCode_LoginErrCode), "用户名不存在")
	}
	user := users[0]
	if user.Password != in.Password {
		return nil, errno.Err(int(pb.ErrCode_LoginErrCode), "用户名或密码错误")
	}
	// todo: 鉴权操作，返回token
	userinfo := pb.UserInfo{
		Id:        uint64(user.ID),
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Phone:     user.Phone,
		AvaUrl:    user.AvaUrl,
		Signature: user.Signature,
	}
	return &pb.LoginResp{Userinfo: &userinfo}, nil
}
