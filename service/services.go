// Package service initializes and manages all service-level dependencies.
package service

import (
	"taas/repository"
	"taas/utils"
)

// Services aggregates all the application services.
type Services struct {
	Tag        *TagService
	Entity     *EntityService
	TagMapping *TagMappingService
	Tenant     *TenantService
}

// NewServices initializes all services with required dependencies.
// It wires repositories and shared clients like cache into service layers.
func NewServices(repos *repository.Repositories, cache *utils.Cache) *Services {
	entityService := NewEntityService(repos.Entity, cache)
	tagService := NewTagService(repos.Tag)
	tenantService := NewTenantService(repos.Tenant)
	tagMappingService := NewTagMappingService(repos.TagMapping, entityService, tagService)

	return &Services{
		Tag:        tagService,
		Entity:     entityService,
		Tenant:     tenantService,
		TagMapping: tagMappingService,
	}
}
