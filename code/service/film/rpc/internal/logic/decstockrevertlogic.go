package logic

import (
	"context"
	"database/sql"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/film/rpc/internal/svc"
	"cinema-ticket/service/film/rpc/types/film"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecStockRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecStockRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecStockRevertLogic {
	return &DecStockRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecStockRevertLogic) DecStockRevert(in *film.DecStockReq) (*film.DecStockReply, error) {
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
		filmInfo.Stock = filmInfo.Stock + 1
		err = l.svcCtx.FilmModel.TxUpdate(tx, filmInfo)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &film.DecStockReply{}, nil
}
