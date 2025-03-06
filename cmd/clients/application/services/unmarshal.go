package services

import (
	"github.com/samircastro27/backend-dashboard/cmd/clients/domain"
	logger "github.com/samircastro27/backend-dashboard/pkg/logger"
)

func UnmarshalOrganizations(orgData []byte) (d []domain.ClientsModel, err error) {
	ro, err := UnmarshalAndMap(orgData, ClientsMapper)
	if err != nil {
	}

	var clients []domain.ClientsModel
	for _, item := range ro.([]interface{}) {
		o, ok := item.(*[]domain.ClientsModel)
		if !ok {
			logger.LogErrWithDetails("", "", "unexpected type: %T", item)
			continue
		}
		clients = append(clients, *o...)
	}
	return clients, nil
}
