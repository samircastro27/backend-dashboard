package services

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/samircastro27/backend-dashboard/cmd/clients/domain"
	logger "github.com/samircastro27/backend-dashboard/pkg/logger"
)

func ClientsMapper(data []interface{}) interface{} {
	oModel := &[]domain.ClientsModel{}
	if len(data) >= 3 {
		v := data[0].([]interface{})
		byteSlice := make([]byte, len(v))
		for i, val := range v {
			if num, ok := val.(float64); ok {
				byteSlice[i] = byte(num)
			}
		}
		id, err := uuid.FromBytes(byteSlice)
		if err != nil {
			logger.LogErrWithDetails("", "", "Error parsing UUID")
		} else {
			appendClients(oModel, id.String(), data)
		}
	}
	return oModel
}

func appendClients(oModel *[]domain.ClientsModel, idString string, data []interface{}) {
	*oModel = append(*oModel, domain.ClientsModel{
		ID:        idString,
		Login:     formatStringOrDefault(data[1]),
		Name:      formatStringOrDefault(data[2]),
		Company:   formatStringOrDefault(data[3]),
		City:      formatStringOrDefault(data[4]),
		Progress:  formatFloatOrDefault(data[5]),
		CreatedAt: formatStringOrDefault(data[6]),
	})
}

// func OrganizationMapper(data []interface{}) interface{} {
// 	oModel := &[]domain.ClientsModel{}
// 	if len(data) >= 3 {
// 		v := data[0].([]interface{})
// 		byteSlice := make([]byte, len(v))
// 		for i, val := range v {
// 			if num, ok := val.(float64); ok {
// 				byteSlice[i] = byte(num)
// 			}
// 		}
// 		id, err := uuid.FromBytes(byteSlice)
// 		if err != nil {
// 			logger.LogErrWithDetails("", "", "Error parsing UUID")
// 		} else {
// 			appendOrganization(oModel, id.String(), data)
// 		}
// 	}
// 	return oModel
// }

// func appendOrganization(oModel *[]domain.ClientsModel, idString string, data []interface{}) {
// 	*oModel = append(*oModel, domain.ClientsModel{
// 		BaseModel: domain.BaseModel{
// 			ID: idString,
// 		},
// 		Name: formatStringOrDefault(data[1]),
// 		BillingInfoModel: domain.BillingInfoModel{
// 			CustomerId:     formatStringOrDefault(data[2]),
// 			SubscriptionId: formatStringOrDefault(data[3]),
// 		},
// 		ProjectModel: domain.ProjectModel{
// 			ProjectId:   formatStringOrDefault(data[4]),
// 			ProjectName: formatStringOrDefault(data[5]),
// 			ProjectSlug: formatStringOrDefault(data[6]),
// 		},
// 		EnvironmentModel: domain.EnvironmentModel{
// 			EnvironmentId:   formatStringOrDefault(data[7]),
// 			EnvironmentName: formatStringOrDefault(data[8]),
// 		},
// 		RuntimeModel: domain.RuntimeModel{
// 			RuntimeId:   formatStringOrDefault(data[9]),
// 			RuntimeName: formatStringOrDefault(data[10]),
// 		},
// 	})
// }

type MapperFunc func([]interface{}) interface{}

func UnmarshalAndMap(data []byte, mapper MapperFunc) (interface{}, error) {
	var rawData [][]interface{}
	err := json.Unmarshal(data, &rawData)
	if err != nil {
		return nil, err
	}

	var result []interface{}
	for _, item := range rawData {
		result = append(result, mapper(item))
	}

	return result, nil
}

func formatStringOrDefault(value interface{}) string {
	if value == nil {
		return ""
	}
	return fmt.Sprintf("%v", value)
}

func formatFloatOrDefault(value interface{}) float64 {
	if value == nil {
		return 0
	}
	return value.(float64)
}
