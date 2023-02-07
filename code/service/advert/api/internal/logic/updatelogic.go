package logic

import (
	"cinema-ticket/common/errorx"
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/advert/api/internal/svc"
	"cinema-ticket/service/advert/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	advertInfo, err := l.svcCtx.AdvertModel.FindOne(l.ctx, req.Id)
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的广告！")
	default:
		return nil, err
	}
	if req.Title != "" {
		advertInfo.Title = req.Title
	}
	if req.Content != "" {
		advertInfo.Content = req.Content
	}
	if req.IsCom != 0 {
		advertInfo.IsCom = req.IsCom
	}
	if req.Status != 0 {
		advertInfo.Status = req.Status
	}
	err = l.svcCtx.AdvertModel.Update(l.ctx, advertInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdateResponse{}, nil
}
