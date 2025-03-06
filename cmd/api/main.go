package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	dapr "github.com/dapr/go-sdk/service/http"
	"github.com/samircastro27/backend-dashboard/cmd/api/middlewares"
	router "github.com/samircastro27/backend-dashboard/cmd/api/routers"
	"github.com/samircastro27/backend-dashboard/pkg/logger"

	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi/v5"
	chimdd "github.com/go-chi/chi/v5/middleware"
)

func main() {
	appPort := "9090"
	if port, ok := os.LookupEnv("APP_API_PORT"); ok {
		appPort = port
	}

	mux := chi.NewRouter()
	mux.Use(chimdd.Logger)
	mux.Use(chimdd.Recoverer)

	mux.Group(func(r chi.Router) {
		r.Use(middlewares.SentryScopeMiddleware())
		r.Use(middlewares.APIKeyMiddleware())
		router.RegisterRouters(r)
	})

	s := dapr.NewServiceWithMux(fmt.Sprintf(":%s", appPort), mux)

	if err := s.AddHealthCheckHandler("/healthz", func(ctx context.Context) error {
		return nil
	}); err != nil {
		log.Fatalf("Error adding handlers: %v", err)
	}

	logger.LogInfoWithDetails("", "", fmt.Sprintf("API service started on port %s", appPort))
	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		sentry.CaptureException(err)
		log.Fatalf("error: %v", err)
	}
}
