package repository

import (
	"context"
	"taas/model"
	"taas/utils"

	"gorm.io/gorm"
)

type TagMappingRepo struct {
	db *gorm.DB
}

func NewTagMappingRepo(db *gorm.DB) *TagMappingRepo {
	return &TagMappingRepo{db: db}
}

func (r *TagMappingRepo) GetTagMappings(ctx context.Context, entityType int, entityID uint) ([]model.TagMapping, error) {
	var tagMappings []model.TagMapping
	err := r.db.WithContext(ctx).Where("entity_type = ? AND entity_id = ?", entityType, entityID).Find(&tagMappings).Error
	return tagMappings, err
}

func (r *TagMappingRepo) GetWithPagination(ctx context.Context,
	entityType int, entityID uint, pagination *utils.Pagination) ([]model.TagMapping, error) {
	var tagMappings []model.TagMapping
	err := r.db.WithContext(ctx).Scopes(pagination.DBscope(r.db)).
		Where("entity_type = ? AND entity_id = ?", entityType, entityID).Order("created_at ASC").Find(&tagMappings).Error
	return tagMappings, err
}

func (r *TagMappingRepo) CreateTagMappings(ctx context.Context, tagMappings []model.TagMapping) error {
	var err error
	for _, mapping := range tagMappings {
		err := r.db.WithContext(ctx).Create(&mapping).Error
		if err != nil {
			return err
		}
	}
	return err
}

func (r *TagMappingRepo) DeleteTagMappings(ctx context.Context, tagIDs []uint, entityID uint, entityType int) error {
	err := r.db.WithContext(ctx).
		Delete(&model.TagMapping{}, "entity_type = ? AND entity_id = ? AND tag_id IN ?", entityType, entityID, tagIDs).Error
	return err
}

func (r *TagMappingRepo) GetCount(ctx context.Context, entityType int, entityID uint) (uint, error) {
	var count int64
	err := r.db.Model(&model.TagMapping{}).WithContext(ctx).Where("entity_type = ? AND entity_id = ?", entityType, entityID).Count(&count).Error
	return uint(count), err
}
