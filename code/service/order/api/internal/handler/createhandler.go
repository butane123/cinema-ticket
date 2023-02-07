package handler

import (
	"net/http"

	"cinema-ticket/common/response"
	"cinema-ticket/service/order/api/internal/logic"
	"cinema-ticket/service/order/api/internal/svc"
	"cinema-ticket/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		response.Response(w, resp, err)
	}
}
