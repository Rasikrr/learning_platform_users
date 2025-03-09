package grpc

import (
	"context"
	pb "github.com/Rasikrr/learning_platform_users/pkg/api/proto/users/grpc"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUsersServer
}

func NewServer(
	server *grpc.Server) {
	pb.RegisterUsersServer(server, &Server{})

}

func (s *Server) GetByEmail(ctx context.Context, in *pb.GetByEmailRequest) (*pb.GetByEmailReply, error) {
	return &pb.GetByEmailReply{
		FirstName: "Rasik",
		LastName:  "Turtulov",
	}, nil
}
