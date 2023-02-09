package logic

import (
	"cinema-ticket/common/errorx"
	"cinema-ticket/common/utils"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/order/api/internal/svc"
	"cinema-ticket/service/order/api/internal/types"

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
	_, err = l.svcCtx.OrderModel.FindOne(l.ctx, req.Id)
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的订单！")
	default:
		return nil, err
	}
	//先更新完数据库，然后删除缓存
	err = l.svcCtx.OrderModel.Delete(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	redisQueryKey := utils.CacheOrderKey + strconv.FormatInt(req.Id, 10)
	_, err = l.svcCtx.RedisClient.Del(redisQueryKey)
	if err != nil {
		return nil, err
	}
	return &types.RemoveResponse{}, nil
}
