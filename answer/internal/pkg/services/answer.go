package services

import (
	"context"

	"github.com/meetmorrowsolonmars/zhopij/answer/internal/domain"
)

type AnswerRepository interface {
	Create(ctx context.Context, answer *domain.Answer) (int64, error)
}

type AnswerService struct {
	repository AnswerRepository
}

func NewAnswerService(repository AnswerRepository) *AnswerService {
	return &AnswerService{
		repository: repository,
	}
}

func (s *AnswerService) Create(ctx context.Context, answer *domain.Answer) (int64, error) {
	return s.repository.Create(ctx, answer)
}
