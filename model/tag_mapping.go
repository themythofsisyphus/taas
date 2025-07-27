package model

import (
	"time"
)

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

type TagMappingRequest struct {
	TagIDs []uint `json:"tag_ids"`
}
