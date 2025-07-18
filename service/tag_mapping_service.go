package service

import (
	"context"
	"taas/model"
	"taas/repository"
	"taas/utils"
)

type TagMappingService struct {
	repo          *repository.TagMappingRepo
	entityService *EntityService
	tagService    *TagService
}

func NewTagMappingService(repo *repository.TagMappingRepo, entityService *EntityService, tagService *TagService) *TagMappingService {
	return &TagMappingService{repo: repo, entityService: entityService, tagService: tagService}
}

func (s *TagMappingService) CreateTagMappings(ctx context.Context,
	reqTags *model.TagMappingRequest, entityType string, entityID uint) ([]model.TagResponse, error) {
	var tagMappings []model.TagMapping

	entity, err := s.entityService.GetEntityByName(ctx, entityType)
	if err != nil {

	}

	for _, tagID := range reqTags.TagIDs {
		tagMapping := model.TagMapping{
			EntityType: entity.ID,
			EntityID:   entityID,
			TagID:      tagID,
		}
		tagMappings = append(tagMappings, tagMapping)
	}
	err = s.repo.CreateTagMappings(ctx, tagMappings)

	if err != nil {

	}

	return s.tagService.GetTagsByIDs(ctx, reqTags.TagIDs)
}

func (s *TagMappingService) RemoveTagMappings(ctx context.Context, reqTags *model.TagMappingRequest, entityType string, entityID uint) error {
	entity, err := s.entityService.GetEntityByName(ctx, entityType)
	if err != nil {

	}

	return s.repo.DeleteTagMappings(ctx, reqTags.TagIDs, entityID, entity.ID)
}

func (s *TagMappingService) GetTagMappings(ctx context.Context, entityType string, entityID uint) ([]model.TagResponse, error) {
	entity, err := s.entityService.GetEntityByName(ctx, entityType)
	if err != nil {

	}
	tagMappings, err := s.repo.GetTagMappings(ctx, entity.ID, entityID)

	if err != nil {
		return nil, err
	}

	var tagIDs []uint
	for _, tm := range tagMappings {
		tagIDs = append(tagIDs, tm.TagID)
	}

	return s.tagService.GetTagsByIDs(ctx, tagIDs)
}

func (s *TagMappingService) GetTagMappingsWithPagination(ctx context.Context,
	entityType string, entityID uint, pagination *utils.Pagination) ([]model.TagResponse, error) {

	entity, err := s.entityService.GetEntityByName(ctx, entityType)
	if err != nil {

	}
	tagMappings, err := s.repo.GetWithPagination(ctx, entity.ID, entityID, pagination)

	if err != nil {
		return nil, err
	}

	var tagIDs []uint
	for _, tm := range tagMappings {
		tagIDs = append(tagIDs, tm.TagID)
	}

	return s.tagService.GetTagsByIDs(ctx, tagIDs)
}
