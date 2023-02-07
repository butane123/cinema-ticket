package logic

import (
	"context"
	"database/sql"

	"github.com/dtm-labs/dtmgrpc"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/order/rpc/internal/svc"
	"cinema-ticket/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRevertLogic {
	return &CreateRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Create的补偿函数
func (l *CreateRevertLogic) CreateRevert(in *order.CreateReq) (*order.CreateReply, error) {
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, err
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, err
	}
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		err = l.svcCtx.OrderModel.TxDelete(tx, in.Id)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &order.CreateReply{Id: in.Id}, nil
}
