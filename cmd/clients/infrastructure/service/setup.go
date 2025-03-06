package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	dapr "github.com/dapr/go-sdk/service/http"
	"github.com/go-chi/chi/v5"
	"github.com/samircastro27/backend-dashboard/pkg/logger"
)

const (
	ERR_HEALTHCHECK_HANDLER = "Error adding handlers: %v"
	API_SERVICE_STARTED     = "API service started on port %s"
	ERR_GENERIC             = "Error: %v"
)

func InitService(port string, routerSetup func(chi.Router)) {
	mux := chi.NewRouter()

	mux.Group(func(r chi.Router) {
		routerSetup(r)
	})

	svc := dapr.NewServiceWithMux(fmt.Sprintf(":%s", port), mux)

	if err := svc.AddHealthCheckHandler("/healthz", func(ctx context.Context) error {
		return nil
	}); err != nil {
		log.Fatalf(ERR_HEALTHCHECK_HANDLER, err)
	}

	logger.LogInfoWithDetails("", "", fmt.Sprintf(API_SERVICE_STARTED, port))
	if err := svc.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf(ERR_GENERIC, err)
	}
}
