package profile

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/meetmorrowsolonmars/zhopij/profile/internal/pb/api/v1/profile"
)

func (i *Implementation) CreateUser(ctx context.Context, request *profile.CreateUserRequest) (*profile.CreateUserResponse, error) {
	_, exists := i.db.FindLogin(request.Login)
	if exists {
		return nil, status.Error(codes.AlreadyExists, "login already exists")
	}

	user := i.db.StoreUser(request.Login)

	return &profile.CreateUserResponse{
		Id: user.Id,
	}, nil

}
