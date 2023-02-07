package logic

import (
	"context"

	"cinema-ticket/service/advert/api/internal/svc"
	"cinema-ticket/service/advert/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NormalListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNormalListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NormalListLogic {
	return &NormalListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NormalListLogic) NormalList() (resp *types.NormalListResponse, err error) {
	listInfo, err := l.svcCtx.AdvertModel.FindAllByIsCom(l.ctx, 0)
	if err != nil {
		return nil, err
	}
	var list []*types.NormalAdvert
	for _, advert := range listInfo {
		list = append(list, &types.NormalAdvert{
			Id:      advert.Id,
			Title:   advert.Title,
			Content: advert.Content,
		})
	}
	return &types.NormalListResponse{List: list}, nil
}
