package logic

import (
	"cinema-ticket/service/order/model"
	"context"
	"database/sql"

	"github.com/dtm-labs/dtmgrpc"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/order/rpc/internal/svc"
	"cinema-ticket/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分布式事务的数据库操作函数，在Rpc层的Create
func (l *CreateLogic) Create(in *order.CreateReq) (*order.CreateReply, error) {
	// 获取 RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, err
	}
	// 获取子事务屏障对象，来源于GRPC
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, err
	}
	var orderId int64
	//开启该子事务屏障
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		sqlRes, err := l.svcCtx.OrderModel.TxInsert(tx, &model.Order{
			Id:     in.Id,
			Uid:    in.Uid,
			Fid:    in.Fid,
			Amount: in.Amount,
			Status: in.Status,
		})
		if err != nil {
			return err
		}
		orderId, err = sqlRes.LastInsertId()
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &order.CreateReply{Id: orderId}, nil
}
