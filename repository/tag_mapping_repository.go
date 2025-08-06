// Package repository provides access to database operations for tag mappings and other domain models.
package repository

import (
	"context"
	"taas/model"
	"taas/utils"

	"gorm.io/gorm"
)

// TagMappingRepo handles CRUD operations for tag mappings between tags and entities.
type TagMappingRepo struct {
	db *gorm.DB
}

// NewTagMappingRepo creates a new instance of TagMappingRepo.
func NewTagMappingRepo(db *gorm.DB) *TagMappingRepo {
	return &TagMappingRepo{db: db}
}

// GetTagMappings fetches all tag mappings for the given entity type and ID.
func (r *TagMappingRepo) GetTagMappings(ctx context.Context, entityType int, entityID uint) ([]model.TagMapping, error) {
	var tagMappings []model.TagMapping
	err := r.db.WithContext(ctx).Where("entity_type = ? AND entity_id = ?", entityType, entityID).Find(&tagMappings).Error
	return tagMappings, err
}

// GetWithPagination returns tag mappings with pagination support for a specific entity type and ID.
func (r *TagMappingRepo) GetWithPagination(ctx context.Context,
	entityType int, entityID uint, pagination *utils.Pagination) ([]model.TagMapping, error) {
	var tagMappings []model.TagMapping
	err := r.db.WithContext(ctx).Scopes(pagination.DBScope(r.db)).
		Where("entity_type = ? AND entity_id = ?", entityType, entityID).
		Order("created_at ASC").
		Find(&tagMappings).Error
	return tagMappings, err
}

// CreateTagMappings inserts multiple tag mappings into the database.
func (r *TagMappingRepo) CreateTagMappings(ctx context.Context, tagMappings []model.TagMapping) error {
	for _, mapping := range tagMappings {
		if err := r.db.WithContext(ctx).Create(&mapping).Error; err != nil {
			return err
		}
	}
	return nil
}

// DeleteTagMappings removes tag mappings that match the given tag IDs, entity type, and entity ID.
func (r *TagMappingRepo) DeleteTagMappings(ctx context.Context, tagIDs []uint, entityID uint, entityType int) error {
	return r.db.WithContext(ctx).
		Delete(&model.TagMapping{}, "entity_type = ? AND entity_id = ? AND tag_id IN ?", entityType, entityID, tagIDs).
		Error
}

// GetCount returns the number of tag mappings for a specific entity type and ID.
func (r *TagMappingRepo) GetCount(ctx context.Context, entityType int, entityID uint) (uint, error) {
	var count int64
	err := r.db.Model(&model.TagMapping{}).
		WithContext(ctx).
		Where("entity_type = ? AND entity_id = ?", entityType, entityID).
		Count(&count).Error
	return uint(count), err
}
