package logic

import (
	"cinema-ticket/common/errorx"
	"cinema-ticket/common/utils"
	"cinema-ticket/service/user/model"
	"context"

	"cinema-ticket/service/user/api/internal/svc"
	"cinema-ticket/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminRegisterLogic {
	return &AdminRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminRegisterLogic) AdminRegister(req *types.AdminRegisterRequest) (resp *types.AdminRegisterResponse, err error) {
	//判断输入邮箱验证码是否正确
	verificationCode, err := l.svcCtx.RedisClient.Get(req.Email)
	if err != nil || verificationCode == "" {
		return nil, errorx.NewCodeError(100, "无发送验证码或验证码已到期！")
	}
	if verificationCode != req.EmailCode {
		return nil, errorx.NewCodeError(100, "输入的验证码不一致！")
	}
	//判断输入手机验证码是否正确
	//判断该手机号是否已被注册
	count, err := l.svcCtx.UserModel.CountByMobile(l.ctx, req.Mobile)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errorx.NewCodeError(100, "该手机号已被注册！")
	}
	//判断该邮箱是否已被注册
	count, err = l.svcCtx.UserModel.CountByEmail(l.ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errorx.NewCodeError(100, "该邮箱已被注册！")
	}
	//插入数据库
	res, err := l.svcCtx.UserModel.InsertWithNewId(l.ctx, &model.User{
		Id:       utils.GenerateNewId(l.svcCtx.RedisClient, "user"),
		Name:     req.Name,
		Gender:   req.Gender,
		Mobile:   req.Mobile,
		Password: utils.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password),
		Email:    req.Email,
		Type:     1,
	})
	if err != nil {
		return nil, err
	}
	userId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &types.AdminRegisterResponse{Id: userId}, nil
	return
}
