package repository

import (
	"github.com/R4yL-dev/pkmc/internal/models"
	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Create(item *models.Item) error {
	return r.db.Create(item).Error
}

func (r *ItemRepository) FindByID(id uint) (*models.Item, error) {
	var item models.Item

	err := r.db.Preload("Extension").Preload("Type").Preload("Language").
		First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}
