package logic

import (
	"cinema-ticket/common/errorx"
	"cinema-ticket/common/kqueue"
	"cinema-ticket/common/utils"
	"cinema-ticket/service/order/api/internal/svc"
	"cinema-ticket/service/order/api/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"

	jsoniter "github.com/json-iterator/go"

	"github.com/go-redis/redis/v8"

	_ "github.com/dtm-labs/driver-gozero"
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

// 使用DTM实现分布式事务，分别为创建订单和删减影票库存
// 使用消息队列存储下单信息
// 利用用分布式锁创建订单
func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	//调用lua脚本，从缓存中判断影票库存是否足够并删减
	client := redis.NewClient(&redis.Options{
		Addr: l.svcCtx.RedisClient.Addr,
	})
	file, err := os.ReadFile("common/scripts/decStock.lua")
	if err != nil {
		return nil, err
	}
	script := redis.NewScript(string(file))
	redisQueryKey := utils.CacheStockKey + strconv.FormatInt(req.Fid, 10)
	res, err := script.Run(context.Background(), client, []string{}, redisQueryKey).Result()
	if err != nil {
		panic(err)
	}
	value := reflect.ValueOf(res).Int()
	//不存在该电影
	if value == 1 {
		return nil, errorx.NewCodeError(100, "查无此Id的电影！")
	}
	//库存不足
	if value == 2 {
		return nil, errorx.NewCodeError(100, "影票库存不足！")
	}
	//把下单信息放到消息队列中
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	newId := utils.GenerateNewId(l.svcCtx.RedisClient, "order")
	err = l.PubKqCreateOrderMessage(newId, userId, req.Fid, req.Amount, req.Status)
	if err != nil {
		return nil, err
	}
	//直接返回订单id
	return &types.CreateResponse{Id: newId}, nil
}

func (l *CreateLogic) PubKqCreateOrderMessage(id, uid, fid, amount, status int64) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonStr, err := json.MarshalToString(kqueue.CreateOrderMessage{
		Id:     id,
		Uid:    uid,
		Fid:    fid,
		Amount: amount,
		Status: status,
	})
	if err != nil {
		return err
	}
	err = l.svcCtx.KqOrderCreateClient.Push(jsonStr)
	if err != nil {
		return err
	}
	return nil
}
