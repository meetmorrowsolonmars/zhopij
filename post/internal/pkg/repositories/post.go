package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/meetmorrowsolonmars/zhopij/post/internal/domain"
)

type PostRepository struct {
	timeout time.Duration
	db      *pgxpool.Pool
}

func NewPostRepository(timeout time.Duration, db *pgxpool.Pool) *PostRepository {
	return &PostRepository{
		timeout: timeout,
		db:      db,
	}
}

func (r *PostRepository) Create(ctx context.Context, post *domain.Post) (int64, error) {
	query := "INSERT INTO post (author_id, title) VALUES ($1, $2) RETURNING id;"

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	row := r.db.QueryRow(ctx, query, post.AuthorId, post.Title)
	err := row.Scan(&post.ID)
	return post.ID, err
}
