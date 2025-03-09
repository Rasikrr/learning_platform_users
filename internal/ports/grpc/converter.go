package grpc

import (
	"github.com/Rasikrr/learning_platform_core/grpc/converters"
	"github.com/Rasikrr/learning_platform_users/internal/domain/entity"
	pb "github.com/Rasikrr/learning_platform_users/pkg/api/proto/users/grpc"
)

func convertGetByEmailResponse(u *entity.User) *pb.GetByEmailResponse {
	return &pb.GetByEmailResponse{
		User: convertUser(u),
	}
}

func convertUser(u *entity.User) *pb.User {
	return &pb.User{
		Id:          u.ID,
		Name:        u.Name,
		LastName:    u.LastName,
		Email:       u.Email,
		Password:    u.Password,
		AccountRole: u.AccountRole.String(),
		CreatedAt:   converters.ConvertToTimestampPb(&u.CreatedAt),
		UpdatedAt:   converters.ConvertToTimestampPb(&u.UpdatedAt),
		DeletedAt:   converters.ConvertToTimestampPb(u.DeletedAt),
	}
}
