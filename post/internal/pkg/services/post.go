package services

import (
	"context"
	"errors"
	"unicode/utf8"

	"github.com/meetmorrowsolonmars/zhopij/post/internal/domain"
)

type PostRepository interface {
	Create(ctx context.Context, post *domain.Post) (int64, error)
}

type PostService struct {
	repository PostRepository
}

func NewPostService(repository PostRepository) *PostService {
	return &PostService{
		repository: repository,
	}
}

func (s *PostService) Create(ctx context.Context, post *domain.Post) (int64, error) {
	if utf8.RuneCountInString(post.Title) > 512 {
		// TODO: use domain error
		return 0, errors.New("invalid post title length")
	}

	return s.repository.Create(ctx, post)
}
