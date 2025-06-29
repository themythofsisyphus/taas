package repository

import (
	"gorm.io/gorm"
	"taas/models"
)

type TagRepository interface {
	GetAll(tenantID uint) ([]models.Tag, error)
	GetByID(id uint, tenantID uint) (models.Tag, error)
	Create(tag *models.Tag) (*models.Tag, error)
	Update(tag *models.Tag) (*models.Tag, error)
	Delete(id uint, tenantID uint) error
}

type repo struct {
	db *gorm.DB
}

func (r *repo) GetAll(tenantID uint) ([]models.Tag, error) {
	var tags []models.Tag
	err := r.db.Find(&tags).Error
	return tags, err
}

func (r *repo) GetByID(id uint, tenantID uint) (models.Tag, error) {
	var tag models.Tag
	err := r.db.First(&tag, id, tenantID).Error
	return tag, err
}

func (r *repo) Create(tag *models.Tag) (*models.Tag, error) {
	err := r.db.Create(tag).Error
	return tag, err
}

func (r *repo) Update(tag *models.Tag) (*models.Tag, error) {
	err := r.db.Save(tag).Error
	return tag, err
}

func (r *repo) Delete(id uint, tenantID uint) error {
	return r.db.Delete(&models.Tag{}, id, tenantID).Error
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &repo{db}
}
