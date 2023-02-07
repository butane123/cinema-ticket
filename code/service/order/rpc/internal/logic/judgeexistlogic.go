package logic

import (
	"cinema-ticket/common/errorx"
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/order/rpc/internal/svc"
	"cinema-ticket/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJudgeExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeExistLogic {
	return &JudgeExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JudgeExistLogic) JudgeExist(in *order.JudgeExistReq) (*order.JudgeExistReply, error) {
	var err error
	_, err = l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
		return &order.JudgeExistReply{}, err
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的订单！")
	default:
		return nil, err
	}
}
