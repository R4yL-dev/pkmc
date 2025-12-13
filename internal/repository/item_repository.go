package repository

import (
	"context"
	"errors"
	"strconv"

	customErr "github.com/R4yL-dev/pkmc/internal/errors"
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
	err := r.db.WithContext(ctx).Create(item).Error
	if err != nil {
		key := "new"
		if item.ID != 0 {
			key = strconv.Itoa(int(item.ID))
		}
		return customErr.NewRepositoryError("create", "item", key, err)
	}
	return nil
}

func (r *itemRepository) FindByID(ctx context.Context, id uint) (*models.Item, error) {
	var item models.Item

	err := r.db.WithContext(ctx).Preload("Extension").Preload("Type").Preload("Language").
		First(&item, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErr.NewRepositoryError("find", "item", strconv.Itoa(int(id)), customErr.ErrEntityNotFound)
		}
		return nil, customErr.NewRepositoryError("find", "item", strconv.Itoa(int(id)), err)
	}
	return &item, nil
}
