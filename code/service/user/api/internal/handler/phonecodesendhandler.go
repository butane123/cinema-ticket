package handler

import (
	"net/http"

	"cinema-ticket/common/response"
	"cinema-ticket/service/user/api/internal/logic"
	"cinema-ticket/service/user/api/internal/svc"
	"cinema-ticket/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PhoneCodeSendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PhoneCodeSendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPhoneCodeSendLogic(r.Context(), svcCtx)
		resp, err := l.PhoneCodeSend(&req)
		response.Response(w, resp, err)
	}
}
