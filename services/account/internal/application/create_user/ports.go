package createuser

import (
	"context"

	"github.com/koneneru/microservices/services/account/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) error
}
