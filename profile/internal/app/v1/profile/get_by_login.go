package profile

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/meetmorrowsolonmars/zhopij/profile/internal/pb/api/v1/profile"
)

func (i *Implementation) GetByLogin(ctx context.Context, request *profile.GetByLoginRequest) (*profile.GetByLoginResponse, error) {
	id, exists := i.db.FindLogin(request.Login)
	if !exists {
		return nil, status.Error(codes.NotFound, "profile not found")
	}

	user, _ := i.db.GetUser(id)

	return &profile.GetByLoginResponse{
		User: &profile.User{
			Id:        user.Id,
			Login:     user.Login,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}, nil
}
