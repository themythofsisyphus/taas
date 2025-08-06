// Package repository provides access to database operations for tags and other domain models.
package repository

import (
	"context"
	"taas/model"
	"taas/utils"

	"gorm.io/gorm"
)

// TagRepository provides access to tag-related database operations.
type TagRepository struct {
	db *gorm.DB
}

// NewTagRepository creates and returns a new TagRepository instance.
func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db}
}

// GetAll retrieves all tags from the database.
func (r *TagRepository) GetAll(ctx context.Context) ([]model.Tag, error) {
	var tags []model.Tag
	err := r.db.WithContext(ctx).Find(&tags).Error
	return tags, err
}

// GetCount returns the total number of tags in the database.
func (r *TagRepository) GetCount(ctx context.Context) (uint, error) {
	var count int64
	err := r.db.Model(&model.Tag{}).WithContext(ctx).Count(&count).Error
	return uint(count), err
}

// GetWithPagination returns a paginated list of tags.
func (r *TagRepository) GetWithPagination(ctx context.Context, pagination *utils.Pagination) ([]model.Tag, error) {
	var tags []model.Tag
	err := r.db.WithContext(ctx).
		Scopes(pagination.DBScope(r.db)).
		Order("id ASC").
		Find(&tags).Error
	return tags, err
}

// GetByID fetches a tag by its unique ID.
func (r *TagRepository) GetByID(ctx context.Context, id uint) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.WithContext(ctx).First(&tag, id).Error
	return &tag, err
}

// GetByIDs fetches multiple tags using their IDs.
func (r *TagRepository) GetByIDs(ctx context.Context, ids []uint) ([]model.Tag, error) {
	var tags []model.Tag
	err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&tags).Error
	return tags, err
}

// Create inserts a new tag into the database.
func (r *TagRepository) Create(ctx context.Context, tag *model.Tag) (*model.Tag, error) {
	err := r.db.WithContext(ctx).Create(tag).Error
	return tag, err
}

// Update modifies an existing tag in the database.
func (r *TagRepository) Update(ctx context.Context, tag *model.Tag) error {
	return r.db.WithContext(ctx).Save(tag).Error
}

// Delete removes a tag from the database by its ID.
func (r *TagRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Tag{}, id).Error
}

// Search searches tags using a full-text query on the `name_tsv` field with pagination.
func (r *TagRepository) Search(ctx context.Context, term string, pagination *utils.Pagination) ([]model.Tag, error) {
	var tags []model.Tag
	err := r.db.WithContext(ctx).
		Raw(`SELECT * FROM tags WHERE name_tsv @@ plainto_tsquery('english', ?)`, term).
		Scopes(pagination.DBScope(r.db)).
		Order("id ASC").
		Scan(&tags).Error
	return tags, err
}

// SearchCount returns the count of tags matching a full-text search query.
func (r *TagRepository) SearchCount(ctx context.Context, term string) (uint, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Raw(`SELECT COUNT(*) FROM tags WHERE name_tsv @@ plainto_tsquery('english', ?)`, term).
		Scan(&count).Error
	return uint(count), err
}
