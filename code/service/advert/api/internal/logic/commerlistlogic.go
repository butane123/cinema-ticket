package logic

import (
	"context"

	"cinema-ticket/service/advert/api/internal/svc"
	"cinema-ticket/service/advert/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommerListLogic {
	return &CommerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommerListLogic) CommerList() (resp *types.CommerListResponse, err error) {
	listInfo, err := l.svcCtx.AdvertModel.FindAllByIsCom(l.ctx, 1)
	if err != nil {
		return nil, err
	}
	var list []*types.CommerAdvert
	for _, advert := range listInfo {
		list = append(list, &types.CommerAdvert{
			Id:      advert.Id,
			Title:   advert.Title,
			Content: advert.Content,
		})
	}
	return &types.CommerListResponse{List: list}, nil
}
