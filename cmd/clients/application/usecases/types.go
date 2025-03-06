package usecases

import (
	drepositories "github.com/samircastro27/backend-dashboard/cmd/clients/domain/repositories"
)

type Dependencies struct {
	Repositories *drepositories.Repositories
}
