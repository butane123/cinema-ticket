package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"cinema-ticket/service/comment/api/internal/svc"
	"cinema-ticket/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List() (resp *types.ListResponse, err error) {
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	listInfo, err := l.svcCtx.CommentModel.FindAllByUid(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	var list []*types.UserComment
	for _, comment := range listInfo {
		list = append(list, &types.UserComment{
			Id:          comment.Id,
			Fid:         comment.Fid,
			Title:       comment.Title,
			Content:     comment.Content,
			IsAnonymous: comment.IsAnonymous,
		})
	}
	return &types.ListResponse{List: list}, nil
}
