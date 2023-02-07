package logic

import (
	"cinema-ticket/common/utils"
	"cinema-ticket/service/film/api/internal/svc"
	"cinema-ticket/service/film/api/internal/types"
	"cinema-ticket/service/film/model"
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
	sqlRes, err := l.svcCtx.FilmModel.InsertWithNewId(l.ctx, &model.Film{
		Id:           utils.GenerateNewId(l.svcCtx.RedisClient, "film"),
		Name:         req.Name,
		Desc:         req.Desc,
		Stock:        req.Stock,
		Amount:       req.Amount,
		Screenwriter: req.Screenwriter,
		Director:     req.Director,
		Length:       req.Length,
		IsSelectSeat: req.IsSelectSeat,
	})
	if err != nil {
		return nil, err
	}
	filmId, err := sqlRes.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &types.CreateResponse{Id: filmId}, nil
}
