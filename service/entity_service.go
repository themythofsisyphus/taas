// Package service provides business logic for managing entities.
package service

import (
	"context"
	"fmt"
	"strconv"
	"taas/model"
	"taas/pkg/tlog"
	"taas/repository"
	"taas/utils"
)

// EntityService handles entity-related operations.
type EntityService struct {
	entityRepo  *repository.EntityRepo
	cacheClient *utils.Cache
}

// NewEntityService creates a new instance of EntityService.
func NewEntityService(repo *repository.EntityRepo, cache *utils.Cache) *EntityService {
	return &EntityService{
		entityRepo:  repo,
		cacheClient: cache,
	}
}

// GetAllEntities retrieves all entities from the repository.
func (s *EntityService) GetAllEntities(ctx context.Context) ([]model.EntityResponse, error) {
	entities, err := s.entityRepo.GetAll(ctx)
	if err != nil {
		tlog.Error("[EntityService] Error retrieving entities: %v", err)
		return nil, err
	}

	responses := make([]model.EntityResponse, len(entities))
	for i, entity := range entities {
		responses[i] = *s.buildEntityResponse(&entity)
	}
	return responses, nil
}

// CreateEntity creates a new entity and stores its ID in the cache.
func (s *EntityService) CreateEntity(ctx context.Context, req *model.EntityRequest) (*model.EntityResponse, error) {
	newEntity := &model.Entity{Name: req.Name}

	created, err := s.entityRepo.Create(ctx, newEntity)
	if err != nil {
		tlog.Error("[EntityService] Error creating entity: %v", err)
		return nil, err
	}

	tenantID, ok := ctx.Value("tenant_id").(uint)
	if !ok {
		return nil, fmt.Errorf("tenant_id missing in context")
	}

	entityKey := utils.EntityCacheKey(created.Name, tenantID)
	if err := s.cacheClient.Set(entityKey, strconv.Itoa(created.ID)); err != nil {
		tlog.Error("[EntityService] Cache save failed for key %s: %v", entityKey, err)
	}

	return s.buildEntityResponse(created), nil
}

// DeleteEntity deletes an entity and removes its cache entry.
func (s *EntityService) DeleteEntity(ctx context.Context, eType string) error {
	err := s.entityRepo.Delete(ctx, eType)
	if err != nil {
		return err
	}

	tenantID, ok := ctx.Value("tenant_id").(uint)
	if !ok {
		return fmt.Errorf("tenant_id missing in context")
	}

	entityKey := utils.EntityCacheKey(eType, tenantID)
	if err := s.cacheClient.Remove(entityKey); err != nil {
		tlog.Error("[EntityService] Cache remove failed for key %s", entityKey)
	}

	return nil
}

// GetEntityByName returns an entity by name, checking cache first.
func (s *EntityService) GetEntityByName(ctx context.Context, name string) (*model.EntityResponse, error) {
	tenantID, ok := ctx.Value("tenant_id").(uint)
	if !ok {
		return nil, fmt.Errorf("tenant_id missing in context")
	}

	entityKey := utils.EntityCacheKey(name, tenantID)
	entityID, err := s.cacheClient.Get(entityKey)
	if err == nil {
		tlog.Info("[EntityService] Cache hit for key %s", entityKey)
		idInt, _ := strconv.Atoi(entityID)
		return &model.EntityResponse{
			ID:   idInt,
			Name: name,
		}, nil
	}

	tlog.Info("[EntityService] Cache miss for key %s", entityKey)

	entity, err := s.entityRepo.GetByName(ctx, name)
	if err != nil {
		tlog.Error("[EntityService] Error retrieving entity by name: %v", err)
		return nil, err
	}
	return s.buildEntityResponse(entity), nil
}

// buildEntityResponse maps model.Entity to model.EntityResponse.
func (s *EntityService) buildEntityResponse(entity *model.Entity) *model.EntityResponse {
	return &model.EntityResponse{
		ID:   entity.ID,
		Name: entity.Name,
	}
}
