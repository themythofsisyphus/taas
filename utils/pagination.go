package utils

import "gorm.io/gorm"

type Pagination struct {
	Page  int
	Limit int
}

func NewPagination(page int, limit int) *Pagination {
	if page < 1 {
		page = 1
	}

	if limit < 1 || limit > 50 {
		limit = 10
	}

	return &Pagination{
		Page:  page,
		Limit: limit,
	}
}

func (p *Pagination) DBscope(db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (p.Page - 1) * p.Limit

		return db.Offset(offset).Limit(p.Limit)
	}
}
