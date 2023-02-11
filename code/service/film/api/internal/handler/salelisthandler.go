package handler

import (
	"cinema-ticket/common/response"
	"net/http"

	"cinema-ticket/service/film/api/internal/logic"
	"cinema-ticket/service/film/api/internal/svc"
)

func SaleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewSaleListLogic(r.Context(), svcCtx)
		resp, err := l.SaleList()
		response.Response(w, resp, err)
	}
}
