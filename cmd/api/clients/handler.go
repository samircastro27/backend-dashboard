package clients

import (
	"fmt"
)

type QueryParams struct {
	OrganizationId string `json:"organizationId"`
}

func buildPath(params *QueryParams, endpoint string) string {
	path := fmt.Sprintf("/v1/%s?%s", endpoint, fmt.Sprintf("organizationId=%s", params.OrganizationId))
	return path
}
