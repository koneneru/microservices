package updateuser

import "github.com/koneneru/microservices/services/account/internal/domain"

type Command struct {
	UserID uint64
	User   domain.User
}
