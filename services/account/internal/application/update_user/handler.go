package updateuser

import "context"

type Handler struct {
	repo UserRepository
}

func New(repo UserRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Handle(ctx context.Context, cmd Command) error {
	return h.repo.DeleteUser(ctx, cmd.UserID, cmd.User)
}
