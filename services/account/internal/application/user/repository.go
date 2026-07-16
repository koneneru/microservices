package user

import (
	"context"

	"github.com/koneneru/microservices/services/account/internal/domain"
)

type Repository interface {
	CreateUser(ctx context.Context, user domain.User) error
	DeleteUser(ctx context.Context, userID uint64) error
	GetUser(ctx context.Context, userID uint64) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset int) ([]domain.User, error)
	UpdateUser(ctx context.Context, userID uint64, user domain.User) error
}
