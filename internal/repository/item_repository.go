package repository

import (
	"context"

	"github.com/R4yL-dev/pkmc/internal/models"
	"gorm.io/gorm"
)

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) Create(ctx context.Context, item *models.Item) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *itemRepository) FindByID(ctx context.Context, id uint) (*models.Item, error) {
	var item models.Item

	err := r.db.WithContext(ctx).Preload("Extension").Preload("Type").Preload("Language").
		First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}
