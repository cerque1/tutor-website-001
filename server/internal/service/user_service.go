package service

import (
	"context"

	"github.com/cerque1/tutor-website-001/internal/dto"
	"github.com/cerque1/tutor-website-001/internal/security"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll(ctx context.Context) (*[]dto.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserService) Create(ctx context.Context, req dto.UserCreate) (dto.User, error) {
	hash, err := security.PasswordHasher(req.Password)
	if err != nil {
		return dto.User{}, err
	}

	req.Password = hash
	return s.repo.Create(ctx, req)
}

func (s *UserService) Get(ctx context.Context, idx uint64) (dto.User, error) {
	return s.repo.Get(ctx, idx)
}

func (s *UserService) Patch(ctx context.Context, idx uint64, req dto.UserPatch) (dto.User, error) {
	return s.repo.Patch(ctx, idx, req)
}

func (s *UserService) Delete(ctx context.Context, idx uint64) error {
	return s.repo.Delete(ctx, idx)
}
