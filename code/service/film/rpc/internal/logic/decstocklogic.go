package logic

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dtm-labs/dtmcli"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/film/rpc/internal/svc"
	"cinema-ticket/service/film/rpc/types/film"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecStockLogic {
	return &DecStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分布式事务的数据库操作函数，在Rpc层的DecStock
func (l *DecStockLogic) DecStock(in *film.DecStockReq) (*film.DecStockReply, error) {
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, err
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, err
	}
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		filmInfo, err := l.svcCtx.FilmModel.FindOne(l.ctx, in.Id)
		if err != nil {
			return err
		}
		//库存不足，走回滚
		if filmInfo.Stock == 0 {
			return dtmcli.ErrFailure
		}
		filmInfo.Stock = filmInfo.Stock - 1
		err = l.svcCtx.FilmModel.TxUpdate(tx, filmInfo)
		if err != nil {
			return err
		}
		return nil
	})
	if err == dtmcli.ErrFailure {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	if err != nil {
		return nil, err
	}
	return &film.DecStockReply{}, nil
}
