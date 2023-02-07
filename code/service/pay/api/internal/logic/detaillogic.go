package logic

import (
	"cinema-ticket/common/errorx"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cinema-ticket/service/pay/api/internal/svc"
	"cinema-ticket/service/pay/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	payInfo, err := l.svcCtx.PayModel.FindOne(l.ctx, req.Id)
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的支付流水！")
	default:
		return nil, err
	}
	return &types.DetailResponse{
		Uid:    userId,
		Oid:    payInfo.Oid,
		Amount: payInfo.Amount,
		Source: payInfo.Source,
		Status: payInfo.Status,
	}, nil
}
