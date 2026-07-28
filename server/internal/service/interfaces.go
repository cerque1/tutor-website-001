package service

import (
	"context"

	"github.com/cerque1/tutor-website-001/internal/dto"
)

type UserRepository interface {
	GetAll(ctx context.Context) (*[]dto.User, error)
	Create(ctx context.Context, req dto.UserCreate) (dto.User, error)
	Get(ctx context.Context, idx uint64) (dto.User, error)
	Patch(ctx context.Context, idx uint64, req dto.UserPatch) (dto.User, error)
	Delete(ctx context.Context, idx uint64) error
}
