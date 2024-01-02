package storage

import (
	"context"
	"github.com/Bubotka/Microservices/geo/pkg/db/adapter"
	"github.com/Bubotka/Microservices/user/domain/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.36.0 --name=UserRepository
type UserRepository interface {
	Create(ctx context.Context, user models.User) error
	GetByUsername(ctx context.Context, username string) (models.User, error)
	CheckUser(ctx context.Context, user models.User) error
	List(ctx context.Context, c adapter.Condition) ([]models.User, error)
}
