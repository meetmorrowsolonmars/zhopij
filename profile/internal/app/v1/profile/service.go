package profile

import "github.com/meetmorrowsolonmars/zhopij/profile/internal/pb/api/v1/profile"

type Implementation struct {
	profile.UnimplementedProfileServiceServer

	db map[int64]string
}

func NewProfileService() *Implementation {
	return &Implementation{
		db: make(map[int64]string),
	}
}
