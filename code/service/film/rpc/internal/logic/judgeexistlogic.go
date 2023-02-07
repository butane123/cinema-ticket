package logic

import (
	"cinema-ticket/common/errorx"
	"cinema-ticket/service/film/rpc/internal/svc"
	"cinema-ticket/service/film/rpc/types/film"
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJudgeExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeExistLogic {
	return &JudgeExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JudgeExistLogic) JudgeExist(in *film.JudgeExistReq) (*film.JudgeExistReply, error) {
	var err error
	_, err = l.svcCtx.FilmModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
		return &film.JudgeExistReply{}, err
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的电影！")
	default:
		return nil, err
	}
}
