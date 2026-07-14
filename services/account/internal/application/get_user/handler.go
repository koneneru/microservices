package getuser

import "context"

type Handler struct {
	repo UserRepository
}

func New(repo UserRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Handle(ctx context.Context, cmd Command) (Result, error) {
	user, err := h.repo.GetUser(ctx, cmd.UserID)

	return Result{User: user}, err
}
