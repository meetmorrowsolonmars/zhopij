package profile

import (
	"google.golang.org/grpc"

	desc "github.com/meetmorrowsolonmars/zhopij/profile/internal/pb/api/v1/profile"
)

type Implementation struct {
	desc.UnimplementedProfileServiceServer

	db map[int64]string
}

func NewProfileServiceImplementation() *Implementation {
	return &Implementation{
		db: make(map[int64]string),
	}
}

func RegisterGRPCServer(s grpc.ServiceRegistrar, srv desc.ProfileServiceServer) {
	desc.RegisterProfileServiceServer(s, srv)
}
