package grpc

import (
	"context"
	pb "github.com/Rasikrr/learning_platform_users/pkg/api/proto/users/grpc"
)

func (s *Server) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.EmptySuccessResponse, error) {
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

func (s *Server) GetByEmail(ctx context.Context, in *pb.GetByEmailRequest) (*pb.GetByEmailResponse, error) {
	user, err := s.usersService.GetByEmail(ctx, in.GetEmail())
	if err != nil {
		return nil, err
	}
	return convertGetByEmailResponse(user), nil
}

func (s *Server) ResetPassword(ctx context.Context, in *pb.ResetPasswordRequest) (*pb.EmptySuccessResponse, error) {
	err := s.usersService.ResetPassword(ctx, in.GetEmail(), in.GetPasswordHash())
	if err != nil {
		return nil, err
	}
	return &pb.EmptySuccessResponse{}, nil
}
