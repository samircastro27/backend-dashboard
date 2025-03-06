package main

import (
	"context"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/samircastro27/backend-dashboard/cmd/clients/application/usecases"
	"github.com/samircastro27/backend-dashboard/cmd/clients/infrastructure"
	"github.com/samircastro27/backend-dashboard/cmd/clients/infrastructure/service"
	"github.com/samircastro27/backend-dashboard/cmd/clients/svc"
	redis "github.com/samircastro27/backend-dashboard/config/redis/client"
	dpc "github.com/samircastro27/backend-dashboard/pkg/dapr"
	"github.com/samircastro27/backend-dashboard/pkg/logger"
)

func main() {
	appPort := "9091"
	if port, ok := os.LookupEnv("APP_LOGS_PORT"); ok {
		appPort = port
	}

	daprClient, err := dpc.GetClient()
	if err != nil {
		logger.LogErrWithDetails("", "", "Error creating Dapr client: %v", err)
		panic(err)
	}

	svcDependencies := &usecases.Dependencies{
		Repositories: svc.NewRepositories(daprClient),
	}
	redis.InitRedis()

	svcCtx := svc.NewServiceContext()

	clients := usecases.NewUsersStruc(context.Background(), svcDependencies, redis.RedisClient, daprClient)

	svcCtx = svcCtx.UsersUseCase(clients)

	service.InitService(appPort, func(r chi.Router) {
		infrastructure.RegisterRouters(r, svcCtx)
	})
}
