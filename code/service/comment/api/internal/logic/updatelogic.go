package logic

import (
	"cinema-ticket/common/errorx"
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/comment/api/internal/svc"
	"cinema-ticket/service/comment/api/internal/types"

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
	commentInfo, err := l.svcCtx.CommentModel.FindOne(l.ctx, req.Id)
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的评论！")
	default:
		return nil, err
	}
	if req.Fid != 0 {
		commentInfo.Fid = req.Fid
	}
	if req.Title != "" {
		commentInfo.Title = req.Title
	}
	if req.Content != "" {
		commentInfo.Content = req.Content
	}
	if req.IsAnonymous != 0 {
		commentInfo.IsAnonymous = req.IsAnonymous
	}
	err = l.svcCtx.CommentModel.Update(l.ctx, commentInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdateResponse{}, nil
}
