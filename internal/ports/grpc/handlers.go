package grpc

import (
	"context"
	pb "github.com/Rasikrr/learning_platform_users/pkg/api/proto/users/grpc"
)

func (s *server) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.EmptySuccessResponse, error) {
	user, err := convertUser(in.GetUser())
	if err != nil {
		return nil, err
	}
	err = s.usersService.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.EmptySuccessResponse{}, nil
}

func (s *server) GetByID(ctx context.Context, in *pb.GetByIDRequest) (*pb.GetByIDResponse, error) {
	user, err := s.usersService.GetByID(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}
	return &pb.GetByIDResponse{User: convertUserToPb(user)}, nil
}

func (s *server) GetByEmail(ctx context.Context, in *pb.GetByEmailRequest) (*pb.GetByEmailResponse, error) {
	user, err := s.usersService.GetByEmail(ctx, in.GetEmail())
	if err != nil {
		return nil, err
	}
	return &pb.GetByEmailResponse{User: convertUserToPb(user)}, nil
}

func (s *server) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.EmptySuccessResponse, error) {
	err := s.usersService.UpdateUser(ctx, convertUpdateUserParamsToEntity(in))
	if err != nil {
		return nil, err
	}
	return &pb.EmptySuccessResponse{}, nil
}

func (s *server) ResetPassword(ctx context.Context, in *pb.ResetPasswordRequest) (*pb.EmptySuccessResponse, error) {
	err := s.usersService.ResetPassword(ctx, in.GetEmail(), in.GetPasswordHash())
	if err != nil {
		return nil, err
	}
	return &pb.EmptySuccessResponse{}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteUserRequest) (*pb.EmptySuccessResponse, error) {
	err := s.usersService.Delete(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}
	return &pb.EmptySuccessResponse{}, nil
}

func (s *server) CheckEnrollment(ctx context.Context, in *pb.CheckEnrollmentRequest) (*pb.CheckEnrollmentResponse, error) {
	enrolled, err := s.enrollmentsService.CheckEnrollment(ctx, in.GetUserId(), in.GetCourseId())
	if err != nil {
		return nil, err
	}
	return &pb.CheckEnrollmentResponse{Enrolled: enrolled}, nil
}

func (s *server) GetUserEnrollments(ctx context.Context, in *pb.GetUserEnrollmentsRequest) (*pb.GetUserEnrollmentsResponse, error) {
	enrollments, err := s.enrollmentsService.GetUserEnrollments(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}
	return &pb.GetUserEnrollmentsResponse{Enrollments: convertEnrollmentsToPb(enrollments)}, nil
}

func (s *server) Enroll(ctx context.Context, in *pb.EnrollRequest) (*pb.EmptySuccessResponse, error) {
	err := s.enrollmentsService.Enroll(ctx, in.GetUserId(), in.GetCourseId())
	if err != nil {
		return nil, err
	}
	return &pb.EmptySuccessResponse{}, nil
}
