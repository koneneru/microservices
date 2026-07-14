package getusers

import "github.com/koneneru/microservices/services/account/internal/domain"

type Result struct {
	Users []domain.User
}
