package handler

import (
	"net/http"

	"cinema-ticket/common/response"
	"cinema-ticket/service/order/api/internal/logic"
	"cinema-ticket/service/order/api/internal/svc"
	"cinema-ticket/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		response.Response(w, resp, err)
	}
}
