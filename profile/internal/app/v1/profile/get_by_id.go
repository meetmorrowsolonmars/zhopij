package profile

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/meetmorrowsolonmars/zhopij/profile/internal/pb/api/v1/profile"
)

func (i *Implementation) GetById(ctx context.Context, request *profile.GetByIdRequest) (*profile.GetByIdResponse, error) {
	user, exists := i.db.GetUser(request.Id)
	if !exists {
		return nil, status.Error(codes.NotFound, "profile not found")
	}

	return &profile.GetByIdResponse{
		User: &profile.User{
			Id:        user.Id,
			Login:     user.Login,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}, nil
}
