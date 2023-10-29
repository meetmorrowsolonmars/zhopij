package quiz

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"

	desc "github.com/meetmorrowsolonmars/zhopij/quiz/internal/pb/api/v1/quiz"
	"github.com/meetmorrowsolonmars/zhopij/quiz/internal/pkg/services"
)

type Implementation struct {
	desc.UnimplementedQuizServiceServer

	service *services.QuizService
	logger  *zap.SugaredLogger
}

func NewQuizServiceImplementation(service *services.QuizService, logger *zap.SugaredLogger) *Implementation {
	return &Implementation{
		service: service,
		logger:  logger,
	}
}

func RegisterGRPCServer(s grpc.ServiceRegistrar, srv desc.QuizServiceServer) {
	desc.RegisterQuizServiceServer(s, srv)
}
