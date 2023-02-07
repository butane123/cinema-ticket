package logic

import (
	"cinema-ticket/common/utils"
	"cinema-ticket/service/advert/api/internal/svc"
	"cinema-ticket/service/advert/api/internal/types"
	"cinema-ticket/service/advert/model"
	"context"

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
	sqlRes, err := l.svcCtx.AdvertModel.InsertWithNewId(l.ctx, &model.Advert{
		Id:      utils.GenerateNewId(l.svcCtx.RedisClient, "advert"),
		Title:   req.Title,
		Content: req.Content,
		IsCom:   req.IsCom,
		Status:  req.Status,
	})
	if err != nil {
		return nil, err
	}
	advertId, err := sqlRes.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &types.CreateResponse{Id: advertId}, nil
}
