package application

import "github.com/samircastro27/backend-dashboard/cmd/clients/domain"

type UseCase interface {
	Execute() ([]domain.ClientsModel, error)
}
