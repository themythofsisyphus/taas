// Package repository provides access to all application repositories, each of which encapsulates database operations.
package repository

import "gorm.io/gorm"

// Repositories aggregates all individual repositories used throughout the application.
type Repositories struct {
	Tag        *TagRepository    // Handles CRUD operations for tags.
	Entity     *EntityRepo       // Handles CRUD operations for entities.
	TagMapping *TagMappingRepo   // Handles tag-to-entity mapping operations.
	Tenant     *TenantRepository // Handles tenant-related data operations.
}

// NewRepositories initializes and returns a new instance of Repositories.
func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Tag:        NewTagRepository(db),
		Entity:     NewEntityRepo(db),
		TagMapping: NewTagMappingRepo(db),
		Tenant:     NewTenantRepository(db),
	}
}
