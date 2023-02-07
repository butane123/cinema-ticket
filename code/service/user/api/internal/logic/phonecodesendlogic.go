package logic

import (
	"context"

	"cinema-ticket/service/user/api/internal/svc"
	"cinema-ticket/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneCodeSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneCodeSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneCodeSendLogic {
	return &PhoneCodeSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// PhoneCodeSendRequest {
// Phone string `json:"phone"`
// }
// PhoneCodeSendResponse {
// }
// 暂未实现，考虑之后使用阿里云短信
func (l *PhoneCodeSendLogic) PhoneCodeSend(req *types.PhoneCodeSendRequest) (resp *types.PhoneCodeSendResponse, err error) {
	// todo: add your logic here and delete this line
	return
}
