package service

import (
	"context"
	"taas/model"
	"taas/repository"
)

type TagService struct {
	repo *repository.TagRepository
}

func NewTagService(repo *repository.TagRepository) *TagService {
	return &TagService{repo}
}

func (s *TagService) GetAllTags(ctx context.Context) ([]model.Tag, error) {
	return s.repo.GetAll(ctx)
}

func (s *TagService) GetTagByID(ctx context.Context, id uint) (model.Tag, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TagService) CreateTag(ctx context.Context, tag *model.Tag) (*model.Tag, error) {
	return s.repo.Create(ctx, tag)
}

func (s *TagService) UpdateTag(ctx context.Context, tag *model.Tag) (*model.Tag, error) {
	return s.repo.Update(ctx, tag)
}

func (s *TagService) DeleteTag(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
