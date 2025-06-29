package models

import (
	"time"
)

type Tag struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null"`
	TenantID  uint   `gorm:"type:bigint;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
