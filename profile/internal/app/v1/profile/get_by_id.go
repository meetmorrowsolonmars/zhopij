package profile

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/meetmorrowsolonmars/zhopij/profile/internal/pb/api/v1/profile"
)

func (i *Implementation) GetById(ctx context.Context, request *profile.GetByIdRequest) (*profile.GetByIdResponse, error) {
	login, exists := i.db[request.Id]
	if !exists {
		return nil, status.Error(codes.NotFound, "profile not found")
	}

	return &profile.GetByIdResponse{
		User: &profile.User{
			Id:        request.Id,
			Login:     login,
			CreatedAt: timestamppb.New(time.Now()),
			UpdatedAt: timestamppb.New(time.Now()),
		},
	}, nil
}
