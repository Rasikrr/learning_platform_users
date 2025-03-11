package grpc

import (
	usersS "github.com/Rasikrr/learning_platform_users/internal/services/users"
	pb "github.com/Rasikrr/learning_platform_users/pkg/api/proto/users/grpc"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUsersServer
	usersService usersS.Service
}

func NewServer(
	srv *grpc.Server,
	usersService usersS.Service,
) {
	s := &server{
		usersService: usersService,
	}
	pb.RegisterUsersServer(srv, s)
}
