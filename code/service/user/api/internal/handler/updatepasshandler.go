package handler

import (
	"cinema-ticket/common/response"
	"net/http"

	"cinema-ticket/service/user/api/internal/logic"
	"cinema-ticket/service/user/api/internal/svc"
	"cinema-ticket/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdatePassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePassRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdatePassLogic(r.Context(), svcCtx)
		resp, err := l.UpdatePass(&req)
		response.Response(w, resp, err)
	}
}
