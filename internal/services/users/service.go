package users

import (
	"context"
	"github.com/Rasikrr/learning_platform_users/internal/domain/entity"
	"github.com/Rasikrr/learning_platform_users/internal/repositories/users"
)

type Service interface {
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}

type service struct {
	usersRepository users.Repository
}

func NewService(
	usersRepository users.Repository,
) Service {
	return &service{
		usersRepository: usersRepository,
	}
}

func (s *service) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	return s.usersRepository.GetByEmail(ctx, email)
}
