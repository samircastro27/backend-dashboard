package repositories

import "context"

type BaseRepository interface {
	FindOneById(ctx context.Context, id string) ([]byte, error)
}

type Repositories struct {
	Users UsersRepository
}
