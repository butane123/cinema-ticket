package logic

import (
	"cinema-ticket/common/errorx"
	"cinema-ticket/common/utils"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/film/api/internal/svc"
	"cinema-ticket/service/film/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req *types.RemoveRequest) (resp *types.RemoveResponse, err error) {
	_, err = l.svcCtx.FilmModel.FindOne(l.ctx, req.Id)
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的电影！")
	default:
		return nil, err
	}
	err = l.svcCtx.FilmModel.Delete(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	redisQueryKey := utils.CacheFilmKey + strconv.FormatInt(req.Id, 10)
	_, err = l.svcCtx.RedisClient.Del(redisQueryKey)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.RedisClient.Del(utils.CacheStockKey + strconv.FormatInt(req.Id, 10))
	if err != nil {
		return nil, err
	}
	return &types.RemoveResponse{}, nil
}
