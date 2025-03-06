package dapr

import (
	"fmt"

	daprcli "github.com/dapr/go-sdk/client"
)

func GetClient() (daprcli.Client, error) {
	client, err := daprcli.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create Dapr client: %v", err)
	}
	return client, nil
}
