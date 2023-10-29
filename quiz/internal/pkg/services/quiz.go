package services

import (
	"context"
	"errors"
	"unicode/utf8"

	"github.com/meetmorrowsolonmars/zhopij/quiz/internal/domain"
)

type QuizRepository interface {
	Create(ctx context.Context, quiz *domain.Quiz) (int64, error)
}

type QuizService struct {
	repository QuizRepository
}

func NewQuizService(repository QuizRepository) *QuizService {
	return &QuizService{
		repository: repository,
	}
}

func (s *QuizService) Create(ctx context.Context, quiz *domain.Quiz) (int64, error) {
	if utf8.RuneCountInString(quiz.Title) > 512 {
		// TODO: use domain error
		return 0, errors.New("invalid quiz title length")
	}

	return s.repository.Create(ctx, quiz)
}
