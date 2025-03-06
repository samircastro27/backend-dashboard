package infrastructure

import (
	"github.com/go-chi/chi/v5"
	"github.com/samircastro27/backend-dashboard/cmd/clients/infrastructure/handlers"
	"github.com/samircastro27/backend-dashboard/cmd/clients/svc"
)

func RegisterRouters(r chi.Router, svcCtx *svc.ServiceContext) {
	r.Get("/v1/clients", handlers.GetUsersHandler(svcCtx))
}
