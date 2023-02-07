package logic

import (
	"cinema-ticket/common/errorx"
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/order/api/internal/svc"
	"cinema-ticket/service/order/api/internal/types"

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
	orderInfo, err := l.svcCtx.OrderModel.FindOne(l.ctx, req.Id)
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的订单！")
	default:
		return nil, err
	}
	if req.Uid != 0 {
		orderInfo.Uid = req.Uid
	}
	if req.Fid != 0 {
		orderInfo.Fid = req.Fid
	}
	if req.Amount != 0 {
		orderInfo.Amount = req.Amount
	}
	if req.Status != 0 {
		orderInfo.Status = req.Status
	}
	err = l.svcCtx.OrderModel.Update(l.ctx, orderInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdateResponse{}, nil
}
