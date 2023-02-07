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

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoginLogic) AdminLogin(req *types.AdminLoginRequest) (resp *types.AdminLoginResponse, err error) {
	//根据mobile判断管理员是否存在
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
	//管理员密码是否正确
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
	return &types.AdminLoginResponse{
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
