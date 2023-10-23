package quiz

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/meetmorrowsolonmars/zhopij/quiz/internal/domain"
	desc "github.com/meetmorrowsolonmars/zhopij/quiz/internal/pb/api/v1/quiz"
)

func (i *Implementation) CreateQuiz(ctx context.Context, request *desc.CreateQuizRequest) (*desc.CreateQuizResponse, error) {
	if request.AuthorId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "author id can't be less than or equal to 0")
	}

	id, err := i.service.Create(ctx, &domain.Quiz{
		AuthorId: request.AuthorId,
		Title:    request.Title,
	})
	if err != nil {
		i.logger.Errorf("quiz creation error: %s", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	i.logger.Infof("quiz created with id %d", id)
	return &desc.CreateQuizResponse{
		Id: id,
	}, nil
}
