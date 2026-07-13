package deleteuser

import "context"

type UserRepository interface {
	DeleteUser(ctx context.Context, userID uint64) error
}
