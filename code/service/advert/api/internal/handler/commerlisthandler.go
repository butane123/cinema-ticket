package handler

import (
	"cinema-ticket/common/response"
	"net/http"

	"cinema-ticket/service/advert/api/internal/logic"
	"cinema-ticket/service/advert/api/internal/svc"
)

func CommerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCommerListLogic(r.Context(), svcCtx)
		resp, err := l.CommerList()
		response.Response(w, resp, err)
	}
}
