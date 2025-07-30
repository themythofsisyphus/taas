package service

import (
	"context"
	"log"
	"strconv"
	"taas/model"
	"taas/repository"
	"taas/utils"
)

type EntityService struct {
	entityRepo  *repository.EntityRepo
	cacheClient *utils.Cache
}

func NewEntityService(repo *repository.EntityRepo, cache *utils.Cache) *EntityService {
	return &EntityService{
		entityRepo:  repo,
		cacheClient: cache,
	}
}

func (s *EntityService) GetAllEntities(ctx context.Context) ([]model.EntityResponse, error) {
	entities, err := s.entityRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	responses := make([]model.EntityResponse, len(entities))
	for indx, entity := range entities {
		responses[indx] = *s.buildEntityResponse(&entity)
	}

	return responses, nil
}

func (s *EntityService) CreateEntity(ctx context.Context, entity *model.EntityRequest) (*model.EntityResponse, error) {
	newEntity := &model.Entity{
		Name: entity.Name,
	}

	createdEntity, err := s.entityRepo.Create(ctx, newEntity)
	if err != nil {
		return nil, err
	}

	entityKey := utils.EntityCacheKey(createdEntity.Name, ctx.Value("tenant_id").(uint))
	s.cacheClient.Set(entityKey, strconv.Itoa(createdEntity.ID))

	return s.buildEntityResponse(createdEntity), nil
}

func (s *EntityService) DeleteEntity(ctx context.Context, eType string) error {
	return s.entityRepo.Delete(ctx, eType)
}

func (s *EntityService) GetEntityByName(ctx context.Context, name string) (*model.EntityResponse, error) {
	entityKey := utils.EntityCacheKey(name, ctx.Value("tenant_id").(uint))
	entityID, err := s.cacheClient.Get(entityKey)

	if err == nil {
		log.Println("[Memcached] Fetched :", entityID)
		idInt, _ := strconv.Atoi(entityID)
		return &model.EntityResponse{
			ID:   idInt,
			Name: name,
		}, nil
	}

	log.Println("[Memcached] Miss :", name)
	entity, err := s.entityRepo.GetByName(ctx, name)

	if err != nil {
		return nil, err
	}
	return s.buildEntityResponse(entity), nil
}

func (s *EntityService) buildEntityResponse(entity *model.Entity) *model.EntityResponse {
	response := &model.EntityResponse{
		ID:   entity.ID,
		Name: entity.Name,
	}

	return response
}
