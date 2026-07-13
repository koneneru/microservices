package createuser

import (
	"context"
	"time"

	"github.com/koneneru/microservices/services/account/internal/domain"
)

type Handler struct {
	repo UserRepository
}

func New(repo UserRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Handle(ctx context.Context, cmd Command) error {
	user := domain.User{
		Login:      cmd.Login,
		Email:      cmd.Email,
		Phone:      cmd.Phone,
		FirstName:  cmd.FirstName,
		LastName:   cmd.LastName,
		MiddleName: cmd.MiddleName,
		Age:        cmd.Age,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}

	return h.repo.CreateUser(ctx, user)
}
