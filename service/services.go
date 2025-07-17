package service

import "taas/repository"

type Services struct {
	Tag        *TagService
	Entity     *EntityService
	TagMapping *TagMappingService
}

func NewServices(repos *repository.Repositories) *Services {
	entityService := NewEntityService(repos.Entity)
	tagService := NewTagService(repos.Tag)
	return &Services{
		Tag:        tagService,
		Entity:     entityService,
		TagMapping: NewTagMappingService(repos.TagMapping, entityService, tagService),
	}
}
