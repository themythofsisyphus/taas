package model

type Entity struct {
	ID       int    `gorm:"primaryKey"`
	TenantID uint   `gorm:"type:bigint;not null;uniqueIndex:uniq_entity"`
	Name     string `gorm:"size:255;not null;uniqueIndex:uniq_entity"`
}

type EntityResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type EntityRequest struct {
	Name string `json:"name"`
}
