package routers

import (
	"github.com/go-chi/chi/v5"
	svc "github.com/samircastro27/backend-dashboard/cmd/api/svc"

	clients "github.com/samircastro27/backend-dashboard/cmd/api/clients"
)

func RegisterRouters(r chi.Router) {
	svcCtx := svc.NewServiceContext()
	RoutersClients(r, svcCtx)
}

func RoutersClients(r chi.Router, svcCtx *svc.ServiceContext) {
	r.Get("/v1/clients", clients.GetUsersHandler(svcCtx))
}
