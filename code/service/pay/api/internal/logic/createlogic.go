package logic

import (
	"cinema-ticket/common/errorx"
	"cinema-ticket/common/utils"
	"cinema-ticket/service/order/rpc/types/order"
	"cinema-ticket/service/pay/api/internal/svc"
	"cinema-ticket/service/pay/api/internal/types"
	"cinema-ticket/service/pay/model"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	//查询订单是否存在
	orderInfo, err := l.svcCtx.OrderRpc.FindById(l.ctx, &order.FindReq{Id: req.Oid})
	if err != nil {
		return nil, err
	}
	//查询订单是否已生成过流水
	_, err = l.svcCtx.PayModel.FindByOid(l.ctx, req.Oid)
	switch err {
	case nil:
		return nil, errorx.NewCodeError(100, "该订单已生成过支付流水！")
	case sqlx.ErrNotFound:
		break
	default:
		return nil, err
	}
	sqlRes, err := l.svcCtx.PayModel.InsertWithNewId(l.ctx, &model.Pay{
		Id:     utils.GenerateNewId(l.svcCtx.RedisClient, "pay"),
		Uid:    userId,
		Oid:    req.Oid,
		Amount: orderInfo.Amount,
		Source: 0,
		Status: 0,
	})
	if err != nil {
		return nil, err
	}
	payId, err := sqlRes.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &types.CreateResponse{Id: payId}, nil
}
