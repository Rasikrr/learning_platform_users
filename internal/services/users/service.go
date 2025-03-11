package users

import (
	"context"
	"github.com/Rasikrr/learning_platform_users/internal/domain/entity"
	"github.com/Rasikrr/learning_platform_users/internal/repositories/users"
)

type Service interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUser(ctx context.Context, params *entity.UpdateUserParams) error
	ResetPassword(ctx context.Context, email, password string) error
	Delete(ctx context.Context, id string) error
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

func (s *service) GetByID(ctx context.Context, id string) (*entity.User, error) {
	return s.usersRepository.GetByID(ctx, id)
}

func (s *service) UpdateUser(ctx context.Context, params *entity.UpdateUserParams) error {
	return s.usersRepository.Update(ctx, params)
}

func (s *service) ResetPassword(ctx context.Context, email, password string) error {
	return s.usersRepository.ResetPassword(ctx, email, password)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.usersRepository.Delete(ctx, id)
}
