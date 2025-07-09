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
}
