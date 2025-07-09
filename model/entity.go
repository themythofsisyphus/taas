package model

type Entity struct {
	ID       uint   `gorm:"primaryKey"`
	TenantID uint   `gorm:"type:bigint;not null;uniqueIndex:uniq_entity"`
	Name     string `gorm:"size:255;not null;uniqueIndex:uniq_entity"`
}
