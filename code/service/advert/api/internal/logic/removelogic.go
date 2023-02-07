package logic

import (
	"cinema-ticket/common/errorx"
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/advert/api/internal/svc"
	"cinema-ticket/service/advert/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req *types.RemoveRequest) (resp *types.RemoveResponse, err error) {
	_, err = l.svcCtx.AdvertModel.FindOne(l.ctx, req.Id)
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的广告！")
	default:
		return nil, err
	}
	err = l.svcCtx.AdvertModel.Delete(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &types.RemoveResponse{}, nil
}
