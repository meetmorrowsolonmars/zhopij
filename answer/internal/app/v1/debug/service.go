package debug

import (
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	desc "github.com/meetmorrowsolonmars/zhopij/answer/internal/pb/api/v1/debug"
)

type AnswerEventProducer interface {
	SendCreateEvent(msg []byte) error
}

type Implementation struct {
	desc.UnimplementedDebugServiceServer

	producer  AnswerEventProducer
	marshaler protojson.MarshalOptions
}

func NewDebugServiceImplementation(producer AnswerEventProducer, marshaler protojson.MarshalOptions) *Implementation {
	return &Implementation{
		producer:  producer,
		marshaler: marshaler,
	}
}

func RegisterGRPCServer(s grpc.ServiceRegistrar, srv desc.DebugServiceServer) {
	desc.RegisterDebugServiceServer(s, srv)
}
