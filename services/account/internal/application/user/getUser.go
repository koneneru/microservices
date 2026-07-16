package user

import (
	"context"

	"github.com/koneneru/microservices/services/account/internal/domain"
)

type GetUserQuery struct {
	UserID uint64
}

type GetUserResult struct {
	User domain.User
}

func (s *Service) Get(ctx context.Context, q GetUserQuery) (GetUserResult, error) {
	user, err := s.repo.GetUser(ctx, q.UserID)

	return GetUserResult{User: user}, err
}
