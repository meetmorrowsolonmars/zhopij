package profile

import (
	"github.com/meetmorrowsolonmars/zhopij/profile/internal/domain"
	"github.com/meetmorrowsolonmars/zhopij/profile/internal/pb/api/v1/profile"
)

type Implementation struct {
	profile.UnimplementedProfileServiceServer

	db *domain.DB
}

func NewProfileService() *Implementation {
	return &Implementation{
		db: domain.NewDB(),
	}
}
