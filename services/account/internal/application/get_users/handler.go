package getusers

import "context"

type Handler struct {
	repo UserRepository
}

func New(repo UserRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) GetUsers(ctx context.Context, cmd Command) (Result, error) {
	users, err := h.repo.GetUsers(ctx, cmd.Limit, cmd.Offset)

	return Result{users: users}, err
}
