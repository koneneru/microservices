package getusers

import (
	"context"

	"github.com/koneneru/microservices/services/account/internal/domain"
)

type UserRepository interface {
	GetUsers(ctx context.Context, limit, offset int) ([]domain.User, error)
}
