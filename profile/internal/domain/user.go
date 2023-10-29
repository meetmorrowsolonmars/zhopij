package domain

import (
	"time"
)

type User struct {
	Id        int64
	Login     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
