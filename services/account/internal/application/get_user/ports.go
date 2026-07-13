package getuser

import (
	"context"

	"github.com/koneneru/microservices/services/account/internal/domain"
)

type UserRepository interface {
	GetUser(ctx context.Context, userID uint64) (domain.User, error)
}
