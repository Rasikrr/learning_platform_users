package grpc

import (
	"context"
	pb "github.com/Rasikrr/learning_platform_users/pkg/api/proto/users/grpc"
)

func (s *Server) GetByEmail(ctx context.Context, in *pb.GetByEmailRequest) (*pb.GetByEmailResponse, error) {
	user, err := s.usersService.GetByEmail(ctx, in.GetEmail())
	if err != nil {
		return nil, err
	}
	return convertGetByEmailResponse(user), nil
}
