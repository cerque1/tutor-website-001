package service

import (
	"context"
	"errors"

	"github.com/cerque1/tutor-website-001/internal/dto"
	"github.com/cerque1/tutor-website-001/internal/middleware"
)

type ReviewService struct {
	repo ReviewRepository
}

func NewReviewService(repo ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) GetAll(
	ctx context.Context,
	limit uint64,
	offset uint64,
) (*[]dto.Review, error) {
	return s.repo.GetAll(ctx, limit, offset)
}

func (s *ReviewService) Create(ctx context.Context, req dto.ReviewCreate) (dto.Review, error) {
	userId, ok := ctx.Value(middleware.UserIDKey).(uint64)
	if !ok {
		return dto.Review{}, errors.New("user id not found in context")
	}
	return s.repo.Create(ctx, req, userId)
}

func (s *ReviewService) Get(ctx context.Context, idx uint64) (dto.Review, error) {
	return s.repo.Get(ctx, idx)
}

func (s *ReviewService) Delete(ctx context.Context, idx uint64) error {
	return s.repo.Delete(ctx, idx)
}
