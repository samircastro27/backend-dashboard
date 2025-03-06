package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/dapr/go-sdk/client"
	drepositories "github.com/samircastro27/backend-dashboard/cmd/clients/domain/repositories"
)

type UsersRepository struct {
	daprCli client.Client
}

func NewUsersRepository(dapr client.Client) drepositories.UsersRepository {
	return &UsersRepository{
		daprCli: dapr,
	}
}

func (o *UsersRepository) FindOneById(ctx context.Context, id string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	queryData := map[string]string{
		"sql":    "SELECT * FROM organizations WHERE id = $1;",
		"params": fmt.Sprintf("[\"%s\"]", id),
	}

	queryResult, err := o.daprCli.InvokeBinding(ctx, &client.InvokeBindingRequest{
		Name:      "postgresdb",
		Operation: "query",
		Metadata:  queryData,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to invoke sql binding: %w", err)
	}

	return queryResult.Data, nil
}

func (o *UsersRepository) FindAllClients(ctx context.Context) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	queryData := map[string]string{
		"sql": "SELECT * FROM clients",
	}

	queryResult, err := o.daprCli.InvokeBinding(ctx, &client.InvokeBindingRequest{
		Name:      "postgresdb",
		Operation: "query",
		Metadata:  queryData,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to invoke sql binding: %w", err)
	}

	return queryResult.Data, nil
}
