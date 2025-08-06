// Package repository provides data access logic for interacting with the database.
package repository

import (
	"context"
	"taas/model"

	"gorm.io/gorm"
)

// EntityRepo handles CRUD operations for the Entity model.
type EntityRepo struct {
	db *gorm.DB
}

// NewEntityRepo creates a new instance of EntityRepo.
func NewEntityRepo(db *gorm.DB) *EntityRepo {
	return &EntityRepo{db}
}

// GetAll retrieves all entities for the current tenant context.
func (r *EntityRepo) GetAll(ctx context.Context) ([]model.Entity, error) {
	var entities []model.Entity
	err := r.db.WithContext(ctx).Find(&entities).Error
	return entities, err
}

// Create inserts a new entity into the database for the current tenant.
func (r *EntityRepo) Create(ctx context.Context, entity *model.Entity) (*model.Entity, error) {
	err := r.db.WithContext(ctx).Create(entity).Error
	return entity, err
}

// Delete removes an entity by its name (eType) within the tenant scope.
func (r *EntityRepo) Delete(ctx context.Context, eType string) error {
	return r.db.WithContext(ctx).Delete(&model.Entity{}, "name = ?", eType).Error
}

// GetByName fetches an entity by name within the tenant context.
func (r *EntityRepo) GetByName(ctx context.Context, name string) (*model.Entity, error) {
	var entity model.Entity
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&entity).Error
	return &entity, err
}
