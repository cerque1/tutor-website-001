package service

import (
	"context"

	"github.com/cerque1/tutor-website-001/internal/model"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll(ctx context.Context) (*[]model.User, error) {
	return s.repo.GetAll(ctx)
}