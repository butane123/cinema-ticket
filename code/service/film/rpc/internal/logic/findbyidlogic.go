package logic

import (
	"context"

	"cinema-ticket/service/film/rpc/internal/svc"
	"cinema-ticket/service/film/rpc/types/film"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByIdLogic {
	return &FindByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByIdLogic) FindById(in *film.FindReq) (*film.FindReply, error) {
	filmInfo, err := l.svcCtx.FilmModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &film.FindReply{Stock: filmInfo.Stock}, nil
}
