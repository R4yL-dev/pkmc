package repository

import (
	"context"

	"github.com/R4yL-dev/pkmc/internal/models"
	"gorm.io/gorm"
)

type extensionRepository struct {
	db *gorm.DB
}

func NewExtensionRepository(db *gorm.DB) ExtensionRepository {
	return &extensionRepository{db: db}
}

func (r *extensionRepository) FindByCode(ctx context.Context, code string) (*models.Extension, error) {
	var ext models.Extension

	err := r.db.WithContext(ctx).Where("code = ?", code).First(&ext).Error
	if err != nil {
		return nil, err
	}
	return &ext, nil
}
