// Package model defines database models and API request/response structures for managing entities and their relationships.
package model

// Entity represents a uniquely named object scoped to a tenant.
type Entity struct {
	ID       int    `gorm:"primaryKey"`
	TenantID uint   `gorm:"type:bigint;not null;uniqueIndex:uniq_entity"`
	Name     string `gorm:"size:255;not null;uniqueIndex:uniq_entity"`

	Tenant Tenant `gorm:"foreignKey:TenantID;constraint:OnDelete:CASCADE;"`
}

// EntityResponse defines the JSON structure returned in API responses for an entity.
type EntityResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// EntityRequest defines the JSON structure expected in API requests to create or update an entity.
type EntityRequest struct {
	Name string `json:"name"`
}
