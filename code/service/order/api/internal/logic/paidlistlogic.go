package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"cinema-ticket/service/order/api/internal/svc"
	"cinema-ticket/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaidListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaidListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaidListLogic {
	return &PaidListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaidListLogic) PaidList() (resp *types.PaidListResponse, err error) {
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	listInfo, err := l.svcCtx.OrderModel.FindAllPaidByUid(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	var list []*types.UserOrder
	for _, order := range listInfo {
		list = append(list, &types.UserOrder{
			Id:     order.Id,
			Uid:    userId,
			Fid:    order.Fid,
			Amount: order.Amount,
			Status: order.Status,
		})
	}
	return &types.PaidListResponse{List: list}, nil
}
