package logic

import (
	"cinema-ticket/common/errorx"
	"cinema-ticket/common/utils"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"cinema-ticket/service/user/api/internal/svc"
	"cinema-ticket/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthLogic {
	return &RefreshAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthLogic) RefreshAuth(Authorization string) (resp *types.RefreshAuthResponse, err error) {
	//获得原token的剩余信息
	restClaims := make(jwt.MapClaims)
	judgeValid, err := jwt.ParseWithClaims(Authorization, restClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(l.svcCtx.Config.Auth.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	//判断是否token有效
	if !judgeValid.Valid {
		return nil, errorx.NewCodeError(100, "Token已失效！")
	}
	//利用过期token的其他值，生成新token等信息
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	userId, err := json.Number(fmt.Sprintf("%v", restClaims["userId"])).Int64()
	if err != nil {
		return nil, err
	}
	jwtToken, err := utils.GenerateJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, userId)
	if err != nil {
		return nil, err
	}

	return &types.RefreshAuthResponse{
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
	return
}
