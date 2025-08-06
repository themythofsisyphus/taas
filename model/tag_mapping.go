// Package model defines database models and request structures related to tag mappings.
package model

import (
	"time"
)

// TagMapping represents the association between a tag and an entity within a tenant's scope.
type TagMapping struct {
	ID         uint `gorm:"primaryKey"`
	TenantID   uint `gorm:"type:bigint;not null;uniqueIndex:uniq_tag_map"`
	TagID      uint `gorm:"type:bigint;not null;uniqueIndex:uniq_tag_map"`
	EntityID   uint `gorm:"type:bigint;not null;uniqueIndex:uniq_tag_map"`
	EntityType int  `gorm:"type:int;not null;uniqueIndex:uniq_tag_map"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Tag    Tag    `gorm:"foreignKey:TagID;constraint:OnDelete:CASCADE;"`
	Tenant Tenant `gorm:"foreignKey:TenantID;constraint:OnDelete:CASCADE;"`
	Entity Entity `gorm:"foreignKey:EntityType;constraint:OnDelete:RESTRICT;"`
}

// TagMappingRequest defines the request payload for mapping tags to an entity.
type TagMappingRequest struct {
	TagIDs []uint `json:"tag_ids"`
}
