package repository

import (
	"context"

	"github.com/R4yL-dev/pkmc/internal/models"
	"gorm.io/gorm"
)

type itemTypeRepository struct {
	db *gorm.DB
}

func NewItemTypeRepository(db *gorm.DB) ItemTypeRepository {
	return &itemTypeRepository{db: db}
}

func (r *itemTypeRepository) FindByName(ctx context.Context, name string) (*models.ItemType, error) {
	var itemType models.ItemType

	err := r.db.WithContext(ctx).Where("name = ?", name).First(&itemType).Error
	if err != nil {
		return nil, err
	}
	return &itemType, nil
}
