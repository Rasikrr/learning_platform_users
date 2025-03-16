package enrollments

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform_users/internal/domain/entity"
	"github.com/Rasikrr/learning_platform_users/internal/repositories/enrollments"
)

type Service interface {
	GetUserEnrollments(ctx context.Context, userID string) ([]*entity.Enrollment, error)
	Enroll(ctx context.Context, userID string, courseID string) error
	CheckEnrollment(ctx context.Context, userID string, courseID string) (bool, error)
}

type service struct {
	enrollmentsRepository enrollments.Repository
}

func NewService(
	enrollmentsRepository enrollments.Repository,
) Service {
	return &service{
		enrollmentsRepository: enrollmentsRepository,
	}
}

func (s *service) Enroll(ctx context.Context, userID string, courseID string) error {
	enrolled, err := s.CheckEnrollment(ctx, userID, courseID)
	if err != nil {
		return err
	}
	if enrolled {
		return errors.New("user already enrolled")
	}
	return s.enrollmentsRepository.Enroll(ctx, userID, courseID)
}

func (s *service) CheckEnrollment(ctx context.Context, userID string, courseID string) (bool, error) {
	exists, err := s.enrollmentsRepository.CheckByUserIDAndCourseID(ctx, userID, courseID)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *service) GetUserEnrollments(ctx context.Context, userID string) ([]*entity.Enrollment, error) {
	enrollments, err := s.enrollmentsRepository.GetUserEnrollments(ctx, userID)
	if err != nil {
		return nil, err
	}
	return enrollments, nil
}
