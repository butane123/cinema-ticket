package logic

import (
	"cinema-ticket/common/utils"
	"cinema-ticket/service/film/rpc/types/film"
	"cinema-ticket/service/order/api/internal/svc"
	"cinema-ticket/service/order/api/internal/types"
	"cinema-ticket/service/order/rpc/types/order"
	"context"
	"encoding/json"
	"fmt"

	_ "github.com/dtm-labs/driver-gozero"
	"github.com/dtm-labs/dtmgrpc"

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
	//查询电影是否存在
	_, err = l.svcCtx.FilmRpc.JudgeExist(l.ctx, &film.JudgeExistReq{Id: req.Fid})
	if err != nil {
		return nil, err
	}
	//下面开始使用DTM实现分布式事务，分别为影票库存减1和订单插入数据库
	//获取电影RPC和订单RPC各自的BuildTarget
	filmRpcBT, err := l.svcCtx.Config.FilmRpc.BuildTarget()
	if err != nil {
		return nil, err
	}
	orderRpcBT, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return nil, err
	}
	//创建saga协议的事务
	dtmServer := l.svcCtx.Config.DtmServer
	gid := dtmgrpc.MustGenGid(dtmServer)
	newId := utils.GenerateNewId(l.svcCtx.RedisClient, "order")
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).Add(orderRpcBT+"/order.order/create", orderRpcBT+"/order.order/createRevert", &order.CreateReq{
		Id:     newId,
		Uid:    userId,
		Fid:    req.Fid,
		Amount: req.Amount,
		Status: req.Status,
	}).Add(filmRpcBT+"/film.film/decStock", filmRpcBT+"/film.film/decStockRevert", &film.DecStockReq{Id: req.Fid})
	//事务提交
	err = saga.Submit()
	if err != nil {
		return nil, err
	}
	return &types.CreateResponse{Id: newId}, nil
}
