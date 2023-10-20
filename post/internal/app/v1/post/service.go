package post

import (
	"github.com/meetmorrowsolonmars/zhopij/post/internal/pkg/services"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	desc "github.com/meetmorrowsolonmars/zhopij/post/internal/pb/api/v1/post"
)

type Implementation struct {
	desc.UnimplementedPostServiceServer

	service *services.PostService
	logger  *zap.SugaredLogger
}

func NewPostServiceImplementation(service *services.PostService, logger *zap.SugaredLogger) *Implementation {
	return &Implementation{
		service: service,
		logger:  logger,
	}
}

func RegisterGRPCServer(s grpc.ServiceRegistrar, srv desc.PostServiceServer) {
	desc.RegisterPostServiceServer(s, srv)
}
