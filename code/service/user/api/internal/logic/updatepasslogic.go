package logic

import (
	"cinema-ticket/common/errorx"
	"cinema-ticket/common/utils"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/user/api/internal/svc"
	"cinema-ticket/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePassLogic {
	return &UpdatePassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePassLogic) UpdatePass(req *types.UpdatePassRequest) (resp *types.UpdatePassResponse, err error) {
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
	userInfo.Password = utils.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password)
	err = l.svcCtx.UserModel.Update(l.ctx, userInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdatePassResponse{}, nil
}
