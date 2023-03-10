package handler

import (
	"net/http"

	"cinema-ticket/common/response"
	"cinema-ticket/service/order/api/internal/logic"
	"cinema-ticket/service/order/api/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List()
		response.Response(w, resp, err)
	}
}
