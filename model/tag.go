// Package model contains database models and request/response structures for tags.
package model

import (
	"time"
)

// Tag represents a tag entity belonging to a specific tenant.
type Tag struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null;uniqueIndex:uniq_tag"`
	TenantID  uint   `gorm:"type:bigint;not null;uniqueIndex:uniq_tag"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Tenant Tenant `gorm:"foreignKey:TenantID;constraint:OnDelete:CASCADE;"`
}

// TagResponse defines the response payload for a tag.
type TagResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"` // also fixed "upated_at" typo
}

// TagRequest defines the request payload to create or update a tag.
type TagRequest struct {
	Name string `json:"name"`
}
