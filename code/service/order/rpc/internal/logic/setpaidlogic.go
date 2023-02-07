package logic

import (
	"context"

	"cinema-ticket/service/order/rpc/internal/svc"
	"cinema-ticket/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetPaidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetPaidLogic {
	return &SetPaidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetPaidLogic) SetPaid(in *order.SetPaidReq) (*order.SetPaidReply, error) {
	orderInfo, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	orderInfo.Status = 1
	err = l.svcCtx.OrderModel.Update(l.ctx, orderInfo)
	if err != nil {
		return nil, err
	}
	return &order.SetPaidReply{}, nil
}
