package service

import (
	"context"
	"log"
	"taas/model"
	"taas/repository"
	"taas/utils"
)

// TagMappingService handles tag mapping logic for entities.
type TagMappingService struct {
	repo          *repository.TagMappingRepo
	entityService *EntityService
	tagService    *TagService
}

// NewTagMappingService creates a new instance of TagMappingService.
func NewTagMappingService(repo *repository.TagMappingRepo, entityService *EntityService, tagService *TagService) *TagMappingService {
	return &TagMappingService{
		repo:          repo,
		entityService: entityService,
		tagService:    tagService,
	}
}

// getEntityID retrieves the internal entity ID by entity type name.
func (s *TagMappingService) getEntityID(ctx context.Context, entityType string) (int, error) {
	entity, err := s.entityService.GetEntityByName(ctx, entityType)
	if err != nil {
		log.Println("[Error][TagMappingService.getEntityID] Failed to get entity:", err)
		return 0, err
	}
	return entity.ID, nil
}

// CreateTagMappings associates multiple tags with a specific entity.
func (s *TagMappingService) CreateTagMappings(ctx context.Context,
	reqTags *model.TagMappingRequest, entityType string, entityID uint) ([]model.TagResponse, error) {

	entityTypeID, err := s.getEntityID(ctx, entityType)
	if err != nil {
		return nil, err
	}

	tagMappings := make([]model.TagMapping, 0, len(reqTags.TagIDs))
	for _, tagID := range reqTags.TagIDs {
		tagMappings = append(tagMappings, model.TagMapping{
			EntityType: entityTypeID,
			EntityID:   entityID,
			TagID:      tagID,
		})
	}

	if err := s.repo.CreateTagMappings(ctx, tagMappings); err != nil {
		log.Println("[Error][TagMappingService.CreateTagMappings] Failed to create:", err)
		return nil, err
	}

	return s.tagService.GetTagsByIDs(ctx, reqTags.TagIDs)
}

// RemoveTagMappings removes a tag with a specific entity.
func (s *TagMappingService) RemoveTagMappings(ctx context.Context, reqTags *model.TagMappingRequest, entityType string, entityID uint) error {
	entityTypeID, err := s.getEntityID(ctx, entityType)
	if err != nil {
		return err
	}

	return s.repo.DeleteTagMappings(ctx, reqTags.TagIDs, entityID, entityTypeID)
}

// GetTagMappings retrieves all tags mapped to a given entity.
func (s *TagMappingService) GetTagMappings(ctx context.Context, entityType string, entityID uint) ([]model.TagResponse, error) {
	entityTypeID, err := s.getEntityID(ctx, entityType)
	if err != nil {
		return nil, err
	}

	tagMappings, err := s.repo.GetTagMappings(ctx, entityTypeID, entityID)
	if err != nil {
		log.Println("[Error][TagMappingService.GetTagMappings] Failed to get:", err)
		return nil, err
	}

	tagIDs := make([]uint, 0, len(tagMappings))
	for _, tm := range tagMappings {
		tagIDs = append(tagIDs, tm.TagID)
	}

	return s.tagService.GetTagsByIDs(ctx, tagIDs)
}

// GetTagMappingsWithPagination retrieves tag mappings for a given entity with pagination support.
func (s *TagMappingService) GetTagMappingsWithPagination(ctx context.Context,
	entityType string, entityID uint, pagination *utils.Pagination) ([]model.TagResponse, error) {

	entityTypeID, err := s.getEntityID(ctx, entityType)
	if err != nil {
		return nil, err
	}

	tagMappings, err := s.repo.GetWithPagination(ctx, entityTypeID, entityID, pagination)
	if err != nil {
		log.Println("[Error][TagMappingService.GetTagMappingsWithPagination] Failed to get:", err)
		return nil, err
	}

	tagIDs := make([]uint, 0, len(tagMappings))
	for _, tm := range tagMappings {
		tagIDs = append(tagIDs, tm.TagID)
	}

	return s.tagService.GetTagsByIDs(ctx, tagIDs)
}

// GetTagMappingsCount returns the total number of tag mappings for a given entity.
func (s *TagMappingService) GetTagMappingsCount(ctx context.Context, entityType string, entityID uint) (uint, error) {
	entityTypeID, err := s.getEntityID(ctx, entityType)
	if err != nil {
		return 0, err
	}

	count, err := s.repo.GetCount(ctx, entityTypeID, entityID)
	if err != nil {
		log.Println("[Error][TagMappingService.GetTagMappingsCount] Failed to get count:", err)
		return 0, err
	}

	return count, nil
}
