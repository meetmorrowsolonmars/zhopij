package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/meetmorrowsolonmars/zhopij/quiz/internal/domain"
)

type QuizRepository struct {
	timeout time.Duration
	db      *pgxpool.Pool
}

func NewQuizRepository(timeout time.Duration, db *pgxpool.Pool) *QuizRepository {
	return &QuizRepository{
		timeout: timeout,
		db:      db,
	}
}

func (r *QuizRepository) Create(ctx context.Context, quiz *domain.Quiz) (int64, error) {
	query := "INSERT INTO quiz (author_id, title) VALUES ($1, $2) RETURNING id;"

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	row := r.db.QueryRow(ctx, query, quiz.AuthorID, quiz.Title)
	err := row.Scan(&quiz.ID)
	return quiz.ID, err
}
