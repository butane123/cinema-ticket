package handler

import (
	"cinema-ticket/common/response"
	"net/http"

	"cinema-ticket/service/pay/api/internal/logic"
	"cinema-ticket/service/pay/api/internal/svc"
	"cinema-ticket/service/pay/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req)
		response.Response(w, resp, err)
	}
}
