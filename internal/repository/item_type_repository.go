package repository

import (
	"github.com/R4yL-dev/pkmc/internal/models"
	"gorm.io/gorm"
)

type ItemTypeRepository struct {
	db *gorm.DB
}

func NewItemTypeRepository(db *gorm.DB) *ItemTypeRepository {
	return &ItemTypeRepository{db: db}
}

func (r *ItemTypeRepository) FindByName(name string) (*models.ItemType, error) {
	var itemType models.ItemType

	err := r.db.Where("name = ?", name).First(&itemType).Error
	if err != nil {
		return nil, err
	}
	return &itemType, nil
}
