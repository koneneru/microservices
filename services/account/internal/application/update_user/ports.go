package updateuser

import (
	"context"

	"github.com/koneneru/microservices/services/account/internal/domain"
)

type UserRepository interface {
	DeleteUser(ctx context.Context, userID uint64, user domain.User) error
}
