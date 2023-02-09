package logic

import (
	"cinema-ticket/common/errorx"
	"cinema-ticket/common/utils"
	"cinema-ticket/service/film/api/internal/svc"
	"cinema-ticket/service/film/api/internal/types"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	filmInfo, err := l.svcCtx.FilmModel.FindOne(l.ctx, req.Id)
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的电影！")
	default:
		return nil, err
	}
	if req.Name != "" {
		filmInfo.Name = req.Name
	}
	if req.Desc != "" {
		filmInfo.Desc = req.Desc
	}
	if req.Stock != 0 {
		filmInfo.Stock = req.Stock
	}
	if req.Amount != 0 {
		filmInfo.Amount = req.Amount
	}
	if req.Screenwriter != "" {
		filmInfo.Screenwriter = req.Screenwriter
	}
	if req.Director != "" {
		filmInfo.Director = req.Director
	}
	if req.Length != 0 {
		filmInfo.Length = req.Length
	}
	if req.IsSelectSeat != 0 {
		filmInfo.IsSelectSeat = req.IsSelectSeat
	}
	err = l.svcCtx.FilmModel.Update(l.ctx, filmInfo)
	if err != nil {
		return nil, err
	}
	redisQueryKey := utils.CacheFilmKey + strconv.FormatInt(req.Id, 10)
	_, err = l.svcCtx.RedisClient.Del(redisQueryKey)
	if err != nil {
		return nil, err
	}
	if req.Stock != 0 {
		err = l.svcCtx.RedisClient.Set(utils.CacheStockKey+strconv.FormatInt(req.Id, 10), strconv.FormatInt(req.Stock, 10))
		if err != nil {
			return nil, err
		}
	}
	return &types.UpdateResponse{}, nil
}
