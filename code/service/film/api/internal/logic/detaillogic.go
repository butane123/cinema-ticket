package logic

import (
	"cinema-ticket/common/errorx"
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/film/api/internal/svc"
	"cinema-ticket/service/film/api/internal/types"

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
	filmInfo, err := l.svcCtx.FilmModel.FindOne(l.ctx, req.Id)
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的电影！")
	default:
		return nil, err
	}
	return &types.DetailResponse{
		Name:         filmInfo.Name,
		Desc:         filmInfo.Desc,
		Stock:        filmInfo.Stock,
		Amount:       filmInfo.Amount,
		Screenwriter: filmInfo.Screenwriter,
		Director:     filmInfo.Director,
		Length:       filmInfo.Length,
		IsSelectSeat: filmInfo.IsSelectSeat,
	}, nil
}
