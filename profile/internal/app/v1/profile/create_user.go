package profile

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/meetmorrowsolonmars/zhopij/profile/internal/pb/api/v1/profile"
)

func (i *Implementation) CreateUser(ctx context.Context, request *profile.CreateUserRequest) (*profile.CreateUserResponse, error) {
	for _, value := range i.db {
		if request.Login == value {
			return nil, status.Error(codes.AlreadyExists, "login already exists")
		}
	}

	id := int64(len(i.db) + 1)
	i.db[id] = request.Login

	return &profile.CreateUserResponse{
		Id: id,
	}, nil

}
