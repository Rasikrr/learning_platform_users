package grpc

import (
	enrollmentsS "github.com/Rasikrr/learning_platform_users/internal/services/enrollments"
	usersS "github.com/Rasikrr/learning_platform_users/internal/services/users"
	pb "github.com/Rasikrr/learning_platform_users/pkg/api/proto/users/grpc"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUsersServer
	usersService       usersS.Service
	enrollmentsService enrollmentsS.Service
}

func NewServer(
	srv *grpc.Server,
	usersService usersS.Service,
	enrollmentsService enrollmentsS.Service,
) {
	s := &server{
		usersService:       usersService,
		enrollmentsService: enrollmentsService,
	}
	pb.RegisterUsersServer(srv, s)
}
