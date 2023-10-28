package debug

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/meetmorrowsolonmars/zhopij/answer/internal/pb/api/v1/debug"
	"github.com/meetmorrowsolonmars/zhopij/answer/internal/pb/api/v1/message"
)

func (i *Implementation) ProduceCreateAnswerEvent(_ context.Context, request *desc.ProduceCreateAnswerEventRequest) (*desc.ProduceCreateAnswerEventResponse, error) {
	msg := &message.V1CreateAnswerEvent{
		Meta: &message.Meta{
			MessageId: uuid.New().String(),
			Version:   message.MessageVersion_V1_0_0,
			Event:     message.EventType_CREATE,
		},
		Event: &message.V1CreateAnswerEvent_Event{
			QuizId: request.QuizId,
			UserId: request.UserId,
		},
	}

	payload, err := i.marshaler.Marshal(msg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't marshal message: %s", err)
	}

	err = i.producer.SendCreateEvent(payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't produce event: %s", err)
	}

	return &desc.ProduceCreateAnswerEventResponse{}, nil
}
