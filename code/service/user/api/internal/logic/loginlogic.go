package logic

import (
	"cinema-ticket/common/errorx"
	"cinema-ticket/common/utils"
	"context"
	"time"

	"cinema-ticket/service/user/api/internal/svc"
	"cinema-ticket/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	//根据mobile判断用户是否存在
	count, err := l.svcCtx.UserModel.CountByMobile(l.ctx, req.Mobile)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errorx.NewCodeError(100, "该手机号码尚未注册过！")
	}
	userInfo, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err != nil {
		return nil, err
	}
	//用户密码是否正确
	nowPass := utils.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password)
	if nowPass != userInfo.Password {
		return nil, errorx.NewCodeError(100, "输入密码不正确！")
	}
	//生成token并返回
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := utils.GenerateJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, userInfo.Id)
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
