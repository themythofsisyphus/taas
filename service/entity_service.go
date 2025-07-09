package service

import (
	"context"
	"taas/model"
	"taas/repository"
)

type EntityService struct {
	repo *repository.EntityRepo
}

func NewEntityService(repo *repository.EntityRepo) *EntityService {
	return &EntityService{repo: repo}
}

func (s *EntityService) GetAllEntities(ctx context.Context) ([]model.Entity, error) {
	return s.repo.GetAll(ctx)
}

func (s *EntityService) CreateEntity(ctx context.Context, entity *model.Entity) (*model.Entity, error) {
	return s.repo.CreateEntity(ctx, entity)
}

func (s *EntityService) DeleteEntity(ctx context.Context, eType string) error {
	return s.repo.DeleteEntity(ctx, eType)
}
