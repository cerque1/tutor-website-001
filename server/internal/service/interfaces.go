package service

import (
	"context"

	"github.com/cerque1/tutor-website-001/internal/dto"
)

type UserRepository interface {
	GetAll(ctx context.Context) (*[]dto.User, error)
	Create(ctx context.Context, req dto.UserCreate) (dto.User, error)
	Get(ctx context.Context, idx uint64) (dto.User, error)
	GetByEmail(ctx context.Context, email string) (dto.AuthUser, error)
	Patch(ctx context.Context, idx uint64, req dto.UserPatch) (dto.User, error)
	Delete(ctx context.Context, idx uint64) error
}

type ServiceRepository interface {
	GetAll(ctx context.Context) (*[]dto.Service, error)
	Create(ctx context.Context, req dto.ServiceCreate) (dto.Service, error)
	Get(ctx context.Context, idx uint64) (dto.Service, error)
	Patch(ctx context.Context, idx uint64, req dto.ServicePatch) (dto.Service, error)
	Delete(ctx context.Context, idx uint64) error
}

type ReviewRepository interface {
	GetAll(ctx context.Context, limit uint64, offset uint64) (*[]dto.Review, error)
	Create(ctx context.Context, req dto.ReviewCreate, userId uint64) (dto.Review, error)
	Get(ctx context.Context, idx uint64) (dto.Review, error)
	Delete(ctx context.Context, idx uint64) error
}
