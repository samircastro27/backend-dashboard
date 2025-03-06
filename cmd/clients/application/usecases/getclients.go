package usecases

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/go-redis/redis/v8"
	"github.com/samircastro27/backend-dashboard/cmd/clients/application/services"
	"github.com/samircastro27/backend-dashboard/cmd/clients/domain"
	"github.com/samircastro27/backend-dashboard/pkg/logger"
)

type UsersStruc struct {
	ctx         context.Context
	svcCtx      *Dependencies
	redisClient *redis.Client
	DaprCli     dapr.Client
}

type QueryParams struct {
	SourceName string `json:"sourceName"`
	Source     string `json:"source"`
	Interval   string `json:"interval"`
	Resource   string `json:"resource"`
	EndTime    string `json:"endTime"`
}
type ResponseStruct struct {
	Data ResponseStructData
}
type ResponseStructData struct {
	MetricsData []*MetricsData `json:"metricsData"`
}
type MetricsData struct {
	ID        string          `json:"id"`
	Namespace string          `json:"namespace"`
	Container string          `json:"container"`
	Pod       string          `json:"pod"`
	Node      string          `json:"node"`
	Values    [][]interface{} `json:"values"`
}

func NewUsersStruc(ctx context.Context, svcCtx *Dependencies, redisClient *redis.Client, daprCli dapr.Client) *UsersStruc {
	return &UsersStruc{
		ctx:         ctx,
		svcCtx:      svcCtx,
		redisClient: redisClient,
		DaprCli:     daprCli,
	}
}

func (e *UsersStruc) Execute() ([]domain.ClientsModel, error) {
	orgsByte, err := e.svcCtx.Repositories.Users.FindAllClients(e.ctx)
	if err != nil {
		logger.LogErr("ERROR: ", err)
		return nil, err
	}

	clients, err := services.UnmarshalOrganizations(orgsByte)
	if err != nil {
		return nil, err
	}

	return clients, nil
}
