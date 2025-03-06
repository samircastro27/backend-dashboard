package svc

import (
	"fmt"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/samircastro27/backend-dashboard/cmd/clients/application"
	repo "github.com/samircastro27/backend-dashboard/cmd/clients/domain/repositories"
	dpc "github.com/samircastro27/backend-dashboard/pkg/dapr"
	"github.com/samircastro27/backend-dashboard/pkg/logger"
)

type ServiceContext struct {
	DaprCli      dapr.Client
	Repositories *repo.Repositories
	UseCase      application.UseCase
}

func NewServiceContext() *ServiceContext {
	daprClient, err := dpc.GetClient()
	if err != nil {
		logger.LogErrWithDetails("", "", fmt.Sprintf("Error creating Dapr client: %v", err))
		panic(err)
	}

	return &ServiceContext{
		DaprCli:      daprClient,
		Repositories: NewRepositories(daprClient),
		UseCase:      nil,
	}
}

func (s *ServiceContext) UsersUseCase(useCase application.UseCase) *ServiceContext {
	s.UseCase = useCase
	return s
}
