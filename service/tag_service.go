package service

import (
	"context"
	"log"
	"taas/model"
	"taas/repository"
	"taas/utils"
)

// TagService handles business logic related to tags.
type TagService struct {
	tagRepo *repository.TagRepository
}

// NewTagService returns a new instance of TagService.
func NewTagService(repo *repository.TagRepository) *TagService {
	return &TagService{
		tagRepo: repo,
	}
}

// GetAllTags retrieves all tags.
func (s *TagService) GetAllTags(ctx context.Context) ([]model.TagResponse, error) {
	tags, err := s.tagRepo.GetAll(ctx)
	if err != nil {
		log.Printf("[Error][TagService.GetAllTags] %v", err)
		return nil, err
	}

	return s.buildTagResponses(tags), nil
}

// GetTagsWithPagination retrieves paginated tags.
func (s *TagService) GetTagsWithPagination(ctx context.Context, pagination *utils.Pagination) ([]model.TagResponse, error) {
	tags, err := s.tagRepo.GetWithPagination(ctx, pagination)
	if err != nil {
		log.Printf("[Error][TagService.GetTagsWithPagination] %v", err)
		return nil, err
	}

	return s.buildTagResponses(tags), nil
}

// GetTagsCount returns the total number of tags.
func (s *TagService) GetTagsCount(ctx context.Context) (uint, error) {
	count, err := s.tagRepo.GetCount(ctx)
	if err != nil {
		log.Printf("[Error][TagService.GetTagsCount] %v", err)
		return 0, err
	}

	return count, nil
}

// GetTagByID fetches a tag by its ID.
func (s *TagService) GetTagByID(ctx context.Context, id uint) (*model.TagResponse, error) {
	tag, err := s.tagRepo.GetByID(ctx, id)
	if err != nil {
		log.Printf("[Error][TagService.GetTagByID] %v", err)
		return nil, err
	}

	return s.buildTagResponse(tag), nil
}

// GetTagsByIDs fetches multiple tags by their IDs.
func (s *TagService) GetTagsByIDs(ctx context.Context, ids []uint) ([]model.TagResponse, error) {
	tags, err := s.tagRepo.GetByIDs(ctx, ids)
	if err != nil {
		log.Printf("[Error][TagService.GetTagsByIDs] %v", err)
		return nil, err
	}

	return s.buildTagResponses(tags), nil
}

// CreateTag creates a new tag.
func (s *TagService) CreateTag(ctx context.Context, req *model.TagRequest) (*model.TagResponse, error) {
	newTag := &model.Tag{
		Name: req.Name,
	}

	createdTag, err := s.tagRepo.Create(ctx, newTag)
	if err != nil {
		log.Printf("[Error][TagService.CreateTag] %v", err)
		return nil, err
	}

	return s.buildTagResponse(createdTag), nil
}

// UpdateTag updates an existing tag.
func (s *TagService) UpdateTag(ctx context.Context, tagID uint, req *model.TagRequest) (*model.TagResponse, error) {
	tag, err := s.tagRepo.GetByID(ctx, tagID)
	if err != nil {
		log.Printf("[Error][TagService.UpdateTag] %v", err)
		return nil, err
	}

	tag.Name = req.Name
	if err := s.tagRepo.Update(ctx, tag); err != nil {
		log.Printf("[Error][TagService.UpdateTag] %v", err)
		return nil, err
	}

	return s.buildTagResponse(tag), nil
}

// DeleteTag deletes a tag by its ID.
func (s *TagService) DeleteTag(ctx context.Context, id uint) error {
	if err := s.tagRepo.Delete(ctx, id); err != nil {
		log.Printf("[Error][TagService.DeleteTag] %v", err)
		return err
	}
	return nil
}

// SearchTags performs full-text search on tag names with pagination.
func (s *TagService) SearchTags(ctx context.Context, term string, pagination *utils.Pagination) ([]model.TagResponse, error) {
	tags, err := s.tagRepo.Search(ctx, term, pagination)
	if err != nil {
		log.Printf("[Error][TagService.SearchTags] %v", err)
		return nil, err
	}

	return s.buildTagResponses(tags), nil
}

// SearchTagsCount returns count of search results for a term.
func (s *TagService) SearchTagsCount(ctx context.Context, term string) (uint, error) {
	count, err := s.tagRepo.SearchCount(ctx, term)
	if err != nil {
		log.Printf("[Error][TagService.SearchTagsCount] %v", err)
		return 0, err
	}

	return count, nil
}

// buildTagResponse converts Tag to TagResponse.
func (s *TagService) buildTagResponse(tag *model.Tag) *model.TagResponse {
	return &model.TagResponse{
		ID:        tag.ID,
		Name:      tag.Name,
		CreatedAt: tag.CreatedAt,
		UpdatedAt: tag.UpdatedAt,
	}
}

// buildTagResponses converts a slice of Tag to a slice of TagResponse.
func (s *TagService) buildTagResponses(tags []model.Tag) []model.TagResponse {
	responses := make([]model.TagResponse, len(tags))
	for i, tag := range tags {
		responses[i] = *s.buildTagResponse(&tag)
	}
	return responses
}
