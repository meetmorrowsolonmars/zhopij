package debug

import (
	"google.golang.org/grpc"

	"github.com/meetmorrowsolonmars/zhopij/answer/internal/domain"
	desc "github.com/meetmorrowsolonmars/zhopij/answer/internal/pb/api/v1/debug"
)

type AnswerEventProducer interface {
	SendCreateEvent(answer *domain.Answer) error
}

type Implementation struct {
	desc.UnimplementedDebugServiceServer

	producer AnswerEventProducer
}

func NewDebugServiceImplementation(producer AnswerEventProducer) *Implementation {
	return &Implementation{
		producer: producer,
	}
}

func RegisterGRPCServer(s grpc.ServiceRegistrar, srv desc.DebugServiceServer) {
	desc.RegisterDebugServiceServer(s, srv)
}
