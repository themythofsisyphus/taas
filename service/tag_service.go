package service

import (
	"context"
	"taas/model"
	"taas/repository"
	"taas/utils"
)

type TagService struct {
	tagRepo *repository.TagRepository
}

func NewTagService(repo *repository.TagRepository) *TagService {
	return &TagService{
		tagRepo: repo,
	}
}

func (s *TagService) GetAllTags(ctx context.Context) ([]model.TagResponse, error) {
	tags, err := s.tagRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	tagResponses := make([]model.TagResponse, len(tags))
	for indx, tag := range tags {
		tagResponses[indx] = *s.buildTagResponse(&tag)
	}
	return tagResponses, nil
}

func (s *TagService) GetTagsWithPagination(ctx context.Context, pagination *utils.Pagination) ([]model.TagResponse, error) {
	tags, err := s.tagRepo.GetWithPagination(ctx, pagination)

	if err != nil {
		return nil, err
	}

	tagResponses := make([]model.TagResponse, len(tags))
	for indx, tag := range tags {
		tagResponses[indx] = *s.buildTagResponse(&tag)
	}
	return tagResponses, nil
}

func (s *TagService) GetTagByID(ctx context.Context, id uint) (*model.TagResponse, error) {
	tag, err := s.tagRepo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return s.buildTagResponse(tag), nil
}

func (s *TagService) GetTagsByIDs(ctx context.Context, ids []uint) ([]model.TagResponse, error) {

	tags, err := s.tagRepo.GetByIDs(ctx, ids)

	if err != nil {
		return nil, err
	}

	tagResponses := make([]model.TagResponse, len(tags))
	for indx, tag := range tags {
		tagResponses[indx] = *s.buildTagResponse(&tag)
	}
	return tagResponses, nil
}

func (s *TagService) CreateTag(ctx context.Context, tag *model.TagRequest) (*model.TagResponse, error) {

	newTag := &model.Tag{
		Name: tag.Name,
	}

	createdTag, err := s.tagRepo.Create(ctx, newTag)
	if err != nil {
		return nil, err
	}
	return s.buildTagResponse(createdTag), nil
}

func (s *TagService) UpdateTag(ctx context.Context, tagID uint, tag *model.TagRequest) (*model.TagResponse, error) {
	updatedTag, err := s.tagRepo.GetByID(ctx, tagID)

	if err != nil {
		return nil, err
	}
	updatedTag.Name = tag.Name

	if err := s.tagRepo.Update(ctx, updatedTag); err != nil {
		return nil, err
	}

	return s.buildTagResponse(updatedTag), nil
}

func (s *TagService) DeleteTag(ctx context.Context, id uint) error {
	return s.tagRepo.Delete(ctx, id)
}

func (s *TagService) buildTagResponse(tag *model.Tag) *model.TagResponse {
	response := &model.TagResponse{
		ID:        tag.ID,
		Name:      tag.Name,
		CreatedAt: tag.CreatedAt,
		UpdatedAt: tag.UpdatedAt,
	}

	return response
}
