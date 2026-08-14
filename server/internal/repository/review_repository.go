package repository

import (
	"context"
	"database/sql"

	"github.com/cerque1/tutor-website-001/internal/dto"
)

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) GetAll(
	ctx context.Context,
	limit uint64,
	offset uint64,
) (*[]dto.Review, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`
		SELECT id, user_id, content, rating FROM reviews
		LIMIT $1 OFFSET $2
		`,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reviews := []dto.Review{}

	for rows.Next() {
		var rev dto.Review
		rows.Scan(
			&rev.ID,
			&rev.UserID,
			&rev.Content,
			&rev.Rating,
		)
		reviews = append(reviews, rev)
	}
	return &reviews, nil
}

func (r *ReviewRepository) Create(
	ctx context.Context,
	req dto.ReviewCreate,
	userId uint64,
) (dto.Review, error) {
	var review dto.Review

	err := r.db.QueryRowContext(
		ctx,
		`
		INSERT INTO reviews(user_id, content, rating)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, content, rating
		`,
		userId,
		req.Content,
		req.Rating,
	).Scan(
		&review.ID,
		&review.UserID,
		&review.Content,
		&review.Rating,
	)
	return review, err
}

func (r *ReviewRepository) Get(
	ctx context.Context,
	idx uint64,
) (dto.Review, error) {
	var review dto.Review

	err := r.db.QueryRowContext(
		ctx,
		`
		SELECT id, user_id, content, rating FROM reviews
		WHERE id = $1
		`,
		idx,
	).Scan(
		&review.ID,
		&review.UserID,
		&review.Content,
		&review.Rating,
	)
	return review, err
}

func (r *ReviewRepository) Delete(
	ctx context.Context,
	idx uint64,
) error {
	result, err := r.db.ExecContext(
		ctx,
		`
		DELETE FROM reviews
		WHERE id = $1
		`,
		idx,
	)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if affected == 0 {
        return sql.ErrNoRows
    }

	return nil
}
