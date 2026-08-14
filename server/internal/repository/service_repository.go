package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/cerque1/tutor-website-001/internal/dto"
)

type ServiceRepository struct {
	db *sql.DB
}

func NewServiceRepository(db *sql.DB) *ServiceRepository {
	return &ServiceRepository{db: db}
}

func (r *ServiceRepository) GetAll(
	ctx context.Context,
) (*[]dto.Service, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, title, description, price FROM services",
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	services := []dto.Service{}

	for rows.Next() {
		var s dto.Service
		rows.Scan(
			&s.ID,
			&s.Title,
			&s.Description,
			&s.Price,
		)
		services = append(services, s)
	}
	return &services, nil
}

func (r *ServiceRepository) Create(
	ctx context.Context,
	req dto.ServiceCreate,
) (dto.Service, error) {
	var serv dto.Service

	err := r.db.QueryRowContext(
		ctx,
		`
		INSERT INTO services(title, description, price)
		VALUES ($1, $2, $3)
		RETURNING id, title, description, price
		`,
		req.Title,
		req.Description,
		req.Price,
	).Scan(
		&serv.ID,
		&serv.Title,
		&serv.Description,
		&serv.Title,
	)
	return serv, err
}

func (r *ServiceRepository) Get(
	ctx context.Context,
	idx uint64,
) (dto.Service, error) {
	var serv dto.Service

	err := r.db.QueryRowContext(
		ctx,
		`
		SELECT id, title, description, price FROM services
		WHERE id=$1
		`,
		idx,
	).Scan(
		&serv.ID,
		&serv.Title,
		&serv.Description,
		&serv.Price,
	)
	return serv, err
}

func (r *ServiceRepository) Patch(
	ctx context.Context,
	idx uint64,
	req dto.ServicePatch,
) (dto.Service, error) {
	query := "UPDATE services SET "
	args := []any{}
	i := 1

	if req.Title != nil {
		query += fmt.Sprintf("title=$%d,", i)
		args = append(args, *req.Title)
		i++
	}
	if req.Description != nil{
		query += fmt.Sprintf("description=$%d,", i)
		args = append(args, *req.Description)
		i++
	}
	if req.Price != nil {
		query += fmt.Sprintf("price=$%d,", i)
		args = append(args, *req.Price)
		i++
	}
	query = strings.TrimSuffix(query, ",")

	query += fmt.Sprintf(
		" WHERE id=$%d RETURNING id, title, description, price",
		i,
	)
	args = append(args, idx)

	var serv dto.Service
	fmt.Println(query)

	err := r.db.QueryRowContext(
		ctx,
		query,
		args...,
	).Scan(
		&serv.ID,
		&serv.Title,
		&serv.Description,
		&serv.Price,
	)
	return serv, err
}

func (r *ServiceRepository) Delete(
	ctx context.Context,
	idx uint64,
) error {
	result, err := r.db.ExecContext(
		ctx,
		`
		DELETE FROM services
		WHERE id=$1
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
