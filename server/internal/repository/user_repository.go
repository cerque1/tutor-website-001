package repository

import (
	"context"
	"database/sql"

	"github.com/cerque1/tutor-website-001/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll(ctx context.Context) (*[]model.User, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT * FROM users`,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}

	for rows.Next() {
		var u model.User
		rows.Scan(
			&u.ID,
			&u.Name,
			&u.PasswordHash,
			&u.IsAdmin,
		)
		users = append(users, u)
	}
	return &users, nil
}
