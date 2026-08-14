package service

import (
	"context"

	"github.com/cerque1/tutor-website-001/internal/dto"
)

type ServiceService struct {
	repo ServiceRepository
}

func NewServiceService(repo ServiceRepository) *ServiceService {
	return &ServiceService{repo: repo}
}

func (s *ServiceService) GetAll(ctx context.Context) (*[]dto.Service, error) {
	return s.repo.GetAll(ctx)
}

func (s *ServiceService) Create(ctx context.Context, req dto.ServiceCreate) (dto.Service, error) {
	return s.repo.Create(ctx, req)
}

func (s *ServiceService) Get(ctx context.Context, idx uint64) (dto.Service, error) {
	return s.repo.Get(ctx, idx)
}

func (s *ServiceService) Patch(
	ctx context.Context,
	idx uint64,
	req dto.ServicePatch,
) (dto.Service, error) {
	return s.repo.Patch(ctx, idx, req)
}

func (s *ServiceService) Delete(ctx context.Context, idx uint64) error {
	return s.repo.Delete(ctx, idx)
}
