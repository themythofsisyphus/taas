package model

type Tenant struct {
	ID uint `gorm:"primaryKey;autoIncrement:false"`
}

type TenantRecord struct {
	TenantID uint `json:"tenant_id"`
}
