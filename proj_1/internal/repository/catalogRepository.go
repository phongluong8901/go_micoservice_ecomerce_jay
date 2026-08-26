package repository

import (
	"gorm.io/gorm"
)

type CatalogRepository interface {
}

type catalogRepository struct {
	db *gorm.DB
}

// constructors
func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepository{
		db: db,
	}
}
