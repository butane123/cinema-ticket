package handler

import (
	"cinema-ticket/common/response"
	"net/http"

	"cinema-ticket/service/comment/api/internal/logic"
	"cinema-ticket/service/comment/api/internal/svc"
	"cinema-ticket/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RemoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRemoveLogic(r.Context(), svcCtx)
		resp, err := l.Remove(&req)
		response.Response(w, resp, err)
	}
}
