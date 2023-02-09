package logic

import (
	"cinema-ticket/common/errorx"
	"cinema-ticket/common/utils"
	"cinema-ticket/service/order/model"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/redis"

	jsoniter "github.com/json-iterator/go"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/order/api/internal/svc"
	"cinema-ticket/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	//先查询缓存有没有该数据
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	redisQueryKey := utils.CacheOrderKey + strconv.FormatInt(req.Id, 10)
	success, err := l.svcCtx.RedisClient.Exists(redisQueryKey)
	if err != nil {
		return nil, err
	}
	//缓存有该数据
	if success {
		redisInfo, err := l.svcCtx.RedisClient.Get(redisQueryKey)
		//判断该数据是否为空值
		if redisInfo == "" {
			return nil, errorx.NewCodeError(100, "查无此Id的订单！")
		}
		//json反序列化成对象
		var orderInfo model.Order
		err = json.UnmarshalFromString(redisInfo, &orderInfo)
		if err != nil {
			return nil, err
		}
		return &types.DetailResponse{
			Uid:    orderInfo.Uid,
			Fid:    orderInfo.Fid,
			Amount: orderInfo.Amount,
			Status: orderInfo.Status,
		}, nil
	}
	//缓存没有该数据，则获取分布式锁后查询数据库
	redisLockKey := redisQueryKey
	redisLock := redis.NewRedisLock(l.svcCtx.RedisClient, redisLockKey)
	redisLock.SetExpire(utils.RedisLockExpireSeconds)
	if ok, err := redisLock.Acquire(); !ok || err != nil {
		return nil, errorx.NewCodeError(100, "当前有其他用户正在进行操作，请稍后重试")
	}
	defer func() {
		recover()
		// 释放锁
		redisLock.Release()
	}()
	orderInfo, err := l.svcCtx.OrderModel.FindOne(l.ctx, req.Id)
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		//查不到该数据，存入空值到缓存中
		l.svcCtx.RedisClient.Setex(redisQueryKey, "", utils.RedisLockExpireSeconds)
		return nil, errorx.NewCodeError(100, "查无此Id的订单！")
	default:
		return nil, err
	}
	//查到该数据，存入json序列化后的对象到缓存中
	jsonStr, err := json.MarshalToString(orderInfo)
	l.svcCtx.RedisClient.Setex(redisQueryKey, jsonStr, utils.RedisLockExpireSeconds)
	return &types.DetailResponse{
		Uid:    orderInfo.Uid,
		Fid:    orderInfo.Fid,
		Amount: orderInfo.Amount,
		Status: orderInfo.Status,
	}, nil
}
