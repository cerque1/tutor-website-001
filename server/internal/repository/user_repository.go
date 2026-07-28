package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/cerque1/tutor-website-001/internal/dto"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll(ctx context.Context) (*[]dto.User, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id, name, is_admin FROM users`,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []dto.User{}

	for rows.Next() {
		var u dto.User
		rows.Scan(
			&u.ID,
			&u.Name,
			&u.IsAdmin,
		)
		users = append(users, u)
	}
	return &users, nil
}

func (r *UserRepository) Create(ctx context.Context, req dto.UserCreate) (dto.User, error) {
	var user dto.User

	err := r.db.QueryRowContext(
		ctx,
		`
		INSERT INTO users(name, password_hash)
		VALUES ($1, $2)
		RETURNING id, name, is_admin
		`,
		req.Name,
		req.Password,
	).Scan(
		&user.ID,
		&user.Name,
		&user.IsAdmin,
	)
	return user, err
}

func (r *UserRepository) Get(ctx context.Context, idx uint64) (dto.User, error) {
	var user dto.User

	err := r.db.QueryRowContext(
		ctx,
		`
		SELECT id, name, is_admin FROM users
		WHERE id=$1
		`,
		idx,
	).Scan(
		&user.ID,
		&user.Name,
		&user.IsAdmin,
	)
	return user, err
}

func (r *UserRepository) Patch(ctx context.Context, idx uint64, req dto.UserPatch) (dto.User, error) {
	query := "UPDATE users SET "
	args := []any{}
	i := 1

	if req.Name != nil {
		query += fmt.Sprintf("name=$%d,", i)
		args = append(args, *req.Name)
		i++
	}
	if req.Password != nil{
		query += fmt.Sprintf("password_hash=%d,", i)
		args = append(args, *req.Password)
		i++
	}
	query = strings.TrimSuffix(query, ",")

	query += fmt.Sprintf(
		" WHERE id=$%d RETURNING id, name, is_admin",
		i,
	)
	args = append(args, idx)

	var user dto.User

	err := r.db.QueryRowContext(
		ctx,
		query,
		args...,
	).Scan(
		&user.ID,
		&user.Name,
		&user.IsAdmin,
	)
	return user, err
}

func (r *UserRepository) Delete(ctx context.Context, idx uint64) error {
	result, err := r.db.ExecContext(
		ctx,
		`
		DELETE FROM users
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
