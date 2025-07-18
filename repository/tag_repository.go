package repository

import (
	"context"
	"taas/model"
	"taas/utils"

	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db}
}

func (r *TagRepository) GetAll(ctx context.Context) ([]model.Tag, error) {
	var tags []model.Tag
	err := r.db.WithContext(ctx).Find(&tags).Error
	return tags, err
}

func (r *TagRepository) GetWithPagination(ctx context.Context, pagination *utils.Pagination) ([]model.Tag, error) {
	var tags []model.Tag
	err := r.db.WithContext(ctx).Scopes(pagination.DBscope(r.db)).Order("id ASC").Find(&tags).Error
	return tags, err
}

func (r *TagRepository) GetByID(ctx context.Context, id uint) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.WithContext(ctx).First(&tag, id).Error
	return &tag, err
}

func (r *TagRepository) GetByIDs(ctx context.Context, ids []uint) ([]model.Tag, error) {
	var tags []model.Tag
	err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&tags).Error
	return tags, err
}

func (r *TagRepository) Create(ctx context.Context, tag *model.Tag) (*model.Tag, error) {
	err := r.db.WithContext(ctx).Create(tag).Error
	return tag, err
}

func (r *TagRepository) Update(ctx context.Context, tag *model.Tag) error {
	err := r.db.WithContext(ctx).Save(tag).Error
	return err
}

func (r *TagRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Tag{}, id).Error
}
