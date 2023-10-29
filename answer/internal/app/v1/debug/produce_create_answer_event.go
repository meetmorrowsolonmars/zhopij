package debug

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/meetmorrowsolonmars/zhopij/answer/internal/domain"
	desc "github.com/meetmorrowsolonmars/zhopij/answer/internal/pb/api/v1/debug"
)

func (i *Implementation) ProduceCreateAnswerEvent(_ context.Context, request *desc.ProduceCreateAnswerEventRequest) (*desc.ProduceCreateAnswerEventResponse, error) {
	err := i.producer.SendCreateEvent(&domain.Answer{
		QuizID:   request.QuizId,
		AuthorID: request.AuthorId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't produce event: %s", err)
	}

	return &desc.ProduceCreateAnswerEventResponse{}, nil
}
