package logic

import (
	"cinema-ticket/common/errorx"
	"cinema-ticket/service/order/rpc/types/order"
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/pay/api/internal/svc"
	"cinema-ticket/service/pay/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CallbackLogic) Callback(req *types.CallbackRequest) (resp *types.CallbackResponse, err error) {
	//判断支付流水是否存在
	payInfo, err := l.svcCtx.PayModel.FindOne(l.ctx, req.Id)
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的支付流水！")
	default:
		return nil, err
	}
	//判断订单是否存在
	_, err = l.svcCtx.OrderRpc.JudgeExist(l.ctx, &order.JudgeExistReq{Id: payInfo.Oid})
	if err != nil {
		return nil, err
	}
	//判断流水金额与回调支付金额是否相同
	if payInfo.Amount != req.Amount {
		return nil, errorx.NewCodeError(100, "支付金额与订单金额不相等！")
	}
	//设置支付流水、订单支付生效
	payInfo.Source = req.Source
	payInfo.Status = req.Status
	err = l.svcCtx.PayModel.Update(l.ctx, payInfo)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.OrderRpc.SetPaid(l.ctx, &order.SetPaidReq{Id: payInfo.Oid})
	if err != nil {
		return nil, err
	}
	return &types.CallbackResponse{}, nil
}
