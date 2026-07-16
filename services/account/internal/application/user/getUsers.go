package user

import (
	"context"

	"github.com/koneneru/microservices/services/account/internal/domain"
)

type GetUsersQuery struct {
	Limit  int
	Offset int
}

type GetUsersResult struct {
	Users []domain.User
}

func (s *Service) List(ctx context.Context, q GetUsersQuery) (GetUsersResult, error) {
	users, err := s.repo.GetUsers(ctx, q.Limit, q.Offset)

	return GetUsersResult{Users: users}, err
}
