package svc

import (
	"fmt"

	daprcli "github.com/dapr/go-sdk/client"
	"github.com/samircastro27/backend-dashboard/config/rabbitmq"
	dpc "github.com/samircastro27/backend-dashboard/pkg/dapr"
	"github.com/samircastro27/backend-dashboard/pkg/logger"
)

type ServiceContext struct {
	DaprCli    daprcli.Client
	BrokerConn rabbitmq.AmqpConnection
}

func NewServiceContext() *ServiceContext {
	daprClient, err := dpc.GetClient()
	if err != nil {
		logger.LogErrWithDetails("", "", fmt.Sprintf("Error creating Dapr client: %v", err))
		panic(err)
	}

	conn, err := rabbitmq.ConnectToRabbitMQ()
	if err != nil {
		logger.LogErrWithDetails("", "", fmt.Sprintf("Error connecting to RabbitMQ: %v", err))
		panic(err)
	}

	return &ServiceContext{
		DaprCli:    daprClient,
		BrokerConn: conn,
	}
}
