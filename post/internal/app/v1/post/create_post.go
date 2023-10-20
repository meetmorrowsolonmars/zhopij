package post

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/meetmorrowsolonmars/zhopij/post/internal/domain"
	desc "github.com/meetmorrowsolonmars/zhopij/post/internal/pb/api/v1/post"
)

func (i *Implementation) CreatePost(ctx context.Context, request *desc.CreatePostRequest) (*desc.CreatePostResponse, error) {
	if request.AuthorId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "author id can't be less than or equal to 0")
	}

	id, err := i.service.Create(ctx, &domain.Post{
		AuthorId: request.AuthorId,
		Title:    request.Title,
	})
	if err != nil {
		i.logger.Errorf("post creation error: %s", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	i.logger.Infof("post created with id %d", id)
	return &desc.CreatePostResponse{
		Id: id,
	}, nil
}
