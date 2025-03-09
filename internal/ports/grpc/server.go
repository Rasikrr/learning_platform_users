package grpc

import (
	usersS "github.com/Rasikrr/learning_platform_users/internal/services/users"
	pb "github.com/Rasikrr/learning_platform_users/pkg/api/proto/users/grpc"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUsersServer
	usersService usersS.Service
}

func NewServer(
	server *grpc.Server,
	usersService usersS.Service,
) {
	srv := &Server{
		usersService: usersService,
	}
	pb.RegisterUsersServer(server, srv)
}
