package types

import (
	"net/url"

	"github.com/samircastro27/backend-dashboard/cmd/clients/domain"
)

func GetParams(values url.Values) (domain.QueryParams, error) {
	params := &domain.QueryParams{}
	params.OrganizationId = values.Get("organizationId")
	return *params, nil
}
