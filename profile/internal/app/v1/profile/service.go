package profile

import (
	"google.golang.org/grpc"

	"github.com/meetmorrowsolonmars/zhopij/profile/internal/domain"
	desc "github.com/meetmorrowsolonmars/zhopij/profile/internal/pb/api/v1/profile"
)

type Implementation struct {
	desc.UnimplementedProfileServiceServer

	db *domain.DB
}

func NewProfileServiceImplementation() *Implementation {
	return &Implementation{
		db: domain.NewDB(),
	}
}

func RegisterGRPCServer(s grpc.ServiceRegistrar, srv desc.ProfileServiceServer) {
	desc.RegisterProfileServiceServer(s, srv)
}
