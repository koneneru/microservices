package user

import (
	"context"

	"github.com/koneneru/microservices/services/account/internal/domain"
)

type UpdateUserInput struct {
	UserID uint64
	User   domain.User
}

func (s *Service) Update(ctx context.Context, in UpdateUserInput) error {
	return s.repo.UpdateUser(ctx, in.UserID, in.User)
}
