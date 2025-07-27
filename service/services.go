package service

import "taas/repository"

type Services struct {
	Tag        *TagService
	Entity     *EntityService
	TagMapping *TagMappingService
	Tenant     *TenantService
}

func NewServices(repos *repository.Repositories) *Services {
	entityService := NewEntityService(repos.Entity)
	tagService := NewTagService(repos.Tag)
	tenantService := NewTenantService(repos.Tenant)
	return &Services{
		Tag:        tagService,
		Entity:     entityService,
		Tenant:     tenantService,
		TagMapping: NewTagMappingService(repos.TagMapping, entityService, tagService),
	}
}
