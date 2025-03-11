package grpc

import (
	coreEnum "github.com/Rasikrr/learning_platform_core/enum"
	"github.com/Rasikrr/learning_platform_core/grpc/converters"
	"github.com/Rasikrr/learning_platform_users/internal/domain/entity"
	pb "github.com/Rasikrr/learning_platform_users/pkg/api/proto/users/grpc"
)

func convertUserToPb(u *entity.User) *pb.User {
	if u == nil {
		return nil
	}
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

func convertUser(user *pb.User) (*entity.User, error) {
	role, err := coreEnum.AccountRoleString(user.GetAccountRole())
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:          user.Id,
		Name:        user.Name,
		LastName:    user.LastName,
		Email:       user.Email,
		Password:    user.Password,
		AccountRole: role,
		CreatedAt:   converters.ConvertToTime(user.CreatedAt),
		UpdatedAt:   converters.ConvertToTime(user.UpdatedAt),
		DeletedAt:   converters.ConvertToTimePtr(user.DeletedAt),
	}, nil
}

func convertUpdateUserParamsToEntity(params *pb.UpdateUserRequest) *entity.UpdateUserParams {
	return &entity.UpdateUserParams{
		ID:       params.UserId,
		Name:     params.Name,
		LastName: params.LastName,
	}
}
