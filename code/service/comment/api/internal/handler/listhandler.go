package handler

import (
	"cinema-ticket/common/response"
	"net/http"

	"cinema-ticket/service/comment/api/internal/logic"
	"cinema-ticket/service/comment/api/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List()
		response.Response(w, resp, err)
	}
}
