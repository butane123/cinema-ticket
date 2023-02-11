package logic

import (
	"cinema-ticket/service/film/api/internal/svc"
	"cinema-ticket/service/film/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaleListLogic {
	return &SaleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaleListLogic) SaleList() (resp *types.SaleListResponse, err error) {
	listInfo, err := l.svcCtx.FilmModel.FindOnSale(l.ctx)
	if err != nil {
		return nil, err
	}
	var list []*types.Film
	for _, film := range listInfo {
		list = append(list, &types.Film{
			Name:         film.Name,
			Desc:         film.Desc,
			Stock:        film.Stock,
			Amount:       film.Amount,
			Screenwriter: film.Screenwriter,
			Director:     film.Director,
			Length:       film.Length,
			IsSelectSeat: film.IsSelectSeat,
		})
	}
	return &types.SaleListResponse{List: list}, nil
}
