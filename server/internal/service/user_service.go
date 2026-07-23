package service

import (
	"context"

	"github.com/cerque1/tutor-website-001/internal/model"
	"github.com/cerque1/tutor-website-001/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll(ctx context.Context) ([]model.User, error) {
	return s.repo.GetAll(ctx)
}