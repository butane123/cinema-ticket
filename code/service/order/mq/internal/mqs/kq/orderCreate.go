package kq

import (
	"cinema-ticket/common/kqueue"
	"cinema-ticket/service/film/rpc/types/film"
	"cinema-ticket/service/order/mq/internal/svc"
	"cinema-ticket/service/order/rpc/types/order"
	"context"

	"github.com/dtm-labs/dtmgrpc"

	jsoniter "github.com/json-iterator/go"
)

type OrderCreateMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCreateMq(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCreateMq {
	return &OrderCreateMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 消费函数Consume，集成了Kafka的go-zero框架会自动识别这个消费函数
func (l *OrderCreateMq) Consume(_, val string) error {
	var message kqueue.CreateOrderMessage
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.UnmarshalFromString(val, &message)
	if err != nil {
		return err
	}
	err = l.execService(message)
	if err != nil {
		return err
	}
	return nil
}

// 执行业务
func (l *OrderCreateMq) execService(message kqueue.CreateOrderMessage) error {
	//获取电影RPC和订单RPC各自的BuildTarget
	filmRpcBT, err := l.svcCtx.Config.FilmRpc.BuildTarget()
	if err != nil {
		return err
	}
	orderRpcBT, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return err
	}
	//创建saga协议的事务
	dtmServer := l.svcCtx.Config.DtmServer
	gid := dtmgrpc.MustGenGid(dtmServer)
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).Add(orderRpcBT+"/order.order/create", orderRpcBT+"/order.order/createRevert", &order.CreateReq{
		Id:     message.Id,
		Uid:    message.Uid,
		Fid:    message.Fid,
		Amount: message.Amount,
		Status: message.Status,
	}).Add(filmRpcBT+"/film.film/decStock", filmRpcBT+"/film.film/decStockRevert", &film.DecStockReq{Id: message.Fid})
	//事务提交
	err = saga.Submit()
	if err != nil {
		return err
	}
	return nil
}
