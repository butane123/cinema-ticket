package logic

import (
	"cinema-ticket/common/errorx"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/user/api/internal/svc"
	"cinema-ticket/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInfoLogic) UpdateInfo(req *types.UpdateInfoRequest) (resp *types.UpdateInfoResponse, err error) {
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的用户！")
	default:
		return nil, err
	}
	if req.Name != "" {
		userInfo.Name = req.Name
	}
	if req.Gender != 0 {
		userInfo.Gender = req.Gender
	}
	if req.Mobile != "" {
		userInfo.Mobile = req.Mobile
	}
	if req.Email != "" {
		userInfo.Email = req.Email
	}
	err = l.svcCtx.UserModel.Update(l.ctx, userInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdateInfoResponse{}, nil
}
