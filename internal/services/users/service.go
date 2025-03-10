package users

import (
	"context"
	"github.com/Rasikrr/learning_platform_users/internal/domain/entity"
	"github.com/Rasikrr/learning_platform_users/internal/repositories/users"
)

type Service interface {
	Create(ctx context.Context, user *entity.User) error
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	ResetPassword(ctx context.Context, email, password string) error
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

func (s *service) Create(ctx context.Context, user *entity.User) error {
	return s.usersRepository.Create(ctx, user)
}

func (s *service) ResetPassword(ctx context.Context, email, password string) error {
	return s.usersRepository.ResetPassword(ctx, email, password)
}
