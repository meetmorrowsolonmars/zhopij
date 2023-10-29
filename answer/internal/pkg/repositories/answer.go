package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/meetmorrowsolonmars/zhopij/answer/internal/domain"
)

type AnswerRepository struct {
	timeout time.Duration
	db      *pgxpool.Pool
}

func NewAnswerRepository(timeout time.Duration, db *pgxpool.Pool) *AnswerRepository {
	return &AnswerRepository{
		timeout: timeout,
		db:      db,
	}
}

func (r *AnswerRepository) Create(ctx context.Context, answer *domain.Answer) (int64, error) {
	query := "INSERT INTO answer (quiz_id, author_id) VALUES ($1, $2) RETURNING id;"

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	row := r.db.QueryRow(ctx, query, answer.QuizID, answer.AuthorID)
	err := row.Scan(&answer.ID)
	return answer.ID, err
}
