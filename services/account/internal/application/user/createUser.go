package user

import (
	"context"
	"time"

	"github.com/koneneru/microservices/services/account/internal/domain"
)

type CreateUserInput struct {
	Login      string
	Email      string
	Phone      string
	FirstName  string
	LastName   string
	MiddleName string
	Age        uint32
}

func (s *Service) Create(ctx context.Context, in CreateUserInput) error {
	user := domain.User{
		Login:      in.Login,
		Email:      in.Email,
		Phone:      in.Phone,
		FirstName:  in.FirstName,
		LastName:   in.LastName,
		MiddleName: in.MiddleName,
		Age:        in.Age,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}

	return s.repo.CreateUser(ctx, user)
}
