package handler

import (
	"cinema-ticket/common/response"
	"net/http"

	"cinema-ticket/service/order/api/internal/logic"
	"cinema-ticket/service/order/api/internal/svc"
)

func PaidListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPaidListLogic(r.Context(), svcCtx)
		resp, err := l.PaidList()
		response.Response(w, resp, err)
	}
}
