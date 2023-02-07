package handler

import (
	"cinema-ticket/common/response"
	"net/http"

	"cinema-ticket/service/advert/api/internal/logic"
	"cinema-ticket/service/advert/api/internal/svc"
)

func NormalListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewNormalListLogic(r.Context(), svcCtx)
		resp, err := l.NormalList()
		response.Response(w, resp, err)
	}
}
