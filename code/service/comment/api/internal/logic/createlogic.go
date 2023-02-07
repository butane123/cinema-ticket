package logic

import (
	"cinema-ticket/common/utils"
	"cinema-ticket/service/comment/api/internal/svc"
	"cinema-ticket/service/comment/api/internal/types"
	"cinema-ticket/service/comment/model"
	"cinema-ticket/service/film/rpc/types/film"
	"context"
	"encoding/json"
	"fmt"

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
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	//查询电影是否存在
	_, err = l.svcCtx.FilmRpc.JudgeExist(l.ctx, &film.JudgeExistReq{Id: req.Fid})
	if err != nil {
		return nil, err
	}
	sqlRes, err := l.svcCtx.CommentModel.InsertWithNewId(l.ctx, &model.Comment{
		Id:          utils.GenerateNewId(l.svcCtx.RedisClient, "comment"),
		Uid:         userId,
		Fid:         req.Fid,
		Title:       req.Title,
		Content:     req.Content,
		IsAnonymous: req.IsAnonymous,
	})
	if err != nil {
		return nil, err
	}
	commentId, err := sqlRes.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &types.CreateResponse{Id: commentId}, nil
}
