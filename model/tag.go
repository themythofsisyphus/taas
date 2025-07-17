package model

import (
	"time"
)

type Tag struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null;uniqueIndex:uniq_tag"`
	TenantID  uint   `gorm:"type:bigint;not null;;uniqueIndex:uniq_tag"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// serialisers

type TagResponse struct {
	ID   			uint   `json:"id"`
	Name 			string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"upated_at"`
}

type TagRequest struct {
	Name 			string `json:"name"`
}