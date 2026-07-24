package service

import (
	"context"

	"github.com/cerque1/tutor-website-001/internal/model"
)

type UserRepository interface {
	GetAll(ctx context.Context) (*[]model.User, error)
}
