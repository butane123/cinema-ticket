package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"cinema-ticket/service/order/api/internal/svc"
	"cinema-ticket/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListRequest) (resp *types.ListResponse, err error) {
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	listInfo, err := l.svcCtx.OrderModel.FindAllByUid(l.ctx, userId)
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
	return &types.ListResponse{List: list}, nil
}
