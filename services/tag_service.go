package services

import (
	"taas/models"
	"taas/repository"
)

type TagService interface {
	GetAllTags(tenantID uint) ([]models.Tag, error)
	GetTagByID(id uint, tenantID uint) (models.Tag, error)
	CreateTag(tag *models.Tag) (*models.Tag, error)
	UpdateTag(tag *models.Tag) (*models.Tag, error)
	DeleteTag(id uint, tenantID uint) error
}

type service struct {
	repo repository.TagRepository
}

func (s *service) GetAllTags(tenantID uint) ([]models.Tag, error) {
	return s.repo.GetAll(tenantID)
}

func (s *service) GetTagByID(id uint, tenantID uint) (models.Tag, error) {
	return s.repo.GetByID(id, tenantID)
}

func (s *service) CreateTag(tag *models.Tag) (*models.Tag, error) {
	return s.repo.Create(tag)
}

func (s *service) UpdateTag(tag *models.Tag) (*models.Tag, error) {
	return s.repo.Update(tag)
}

func (s *service) DeleteTag(id uint, tenantID uint) error {
	return s.repo.Delete(id, tenantID)
}

func NewTagService(repo repository.TagRepository) TagService {
	return &service{repo}
}
