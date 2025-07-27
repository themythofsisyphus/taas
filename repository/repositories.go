package repository

import "gorm.io/gorm"

type Repositories struct {
	Tag        *TagRepository
	Entity     *EntityRepo
	TagMapping *TagMappingRepo
	Tenant     *TenantRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Tag:        NewTagRepository(db),
		Entity:     NewEntityRepo(db),
		TagMapping: NewTagMappingRepo(db),
		Tenant:     NewTenantRepository(db),
	}
}
