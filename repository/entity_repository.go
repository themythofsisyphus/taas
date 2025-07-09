package repository

import (
	"context"
	"taas/model"

	"gorm.io/gorm"
)

type EntityRepo struct {
	db *gorm.DB
}

func NewEntityRepo(db *gorm.DB) *EntityRepo {
	return &EntityRepo{db}
}

func (r *EntityRepo) GetAll(ctx context.Context) ([]model.Entity, error) {
	var entities []model.Entity
	err := r.db.WithContext(ctx).Find(&entities).Error
	return entities, err
}

func (r *EntityRepo) CreateEntity(ctx context.Context, entity *model.Entity) (*model.Entity, error) {
	err := r.db.WithContext(ctx).Create(entity).Error
	return entity, err
}

func (r *EntityRepo) DeleteEntity(ctx context.Context, eType string) error {
	return r.db.WithContext(ctx).Delete(&model.Entity{}, "name = ?", eType).Error
}
