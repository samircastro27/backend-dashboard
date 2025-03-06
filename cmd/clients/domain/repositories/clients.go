package repositories

import "context"

type UsersRepository interface {
	BaseRepository
	FindAllClients(ctx context.Context) ([]byte, error)

}
