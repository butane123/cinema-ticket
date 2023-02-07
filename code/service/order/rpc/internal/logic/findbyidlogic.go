package logic

import (
	"context"

	"cinema-ticket/service/order/rpc/internal/svc"
	"cinema-ticket/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByIdLogic {
	return &FindByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByIdLogic) FindById(in *order.FindReq) (*order.FindReply, error) {
	orderInfo, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &order.FindReply{Amount: orderInfo.Amount}, nil
}
