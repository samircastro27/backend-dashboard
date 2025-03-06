package svc

import (
	"github.com/dapr/go-sdk/client"
	repo "github.com/samircastro27/backend-dashboard/cmd/clients/domain/repositories"
	"github.com/samircastro27/backend-dashboard/cmd/clients/infrastructure/repositories"
)

func NewRepositories(daprCli client.Client) *repo.Repositories {
	return &repo.Repositories{
		Users: repositories.NewUsersRepository(daprCli),
	}
}
