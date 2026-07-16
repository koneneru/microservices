package user

import "context"

type DeleteUserInput struct {
	UserID uint64
}

func (s *Service) Delete(ctx context.Context, in DeleteUserInput) error {
	return s.repo.DeleteUser(ctx, in.UserID)
}
