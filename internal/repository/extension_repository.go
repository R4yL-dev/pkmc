package repository

import (
	"context"

	customErr "github.com/R4yL-dev/pkmc/internal/errors"
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

	err := r.db.WithContext(ctx).Preload("Block").Where("code = ?", code).First(&ext).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, customErr.NewRepositoryError("find", "extension", code, customErr.ErrEntityNotFound)
		}
		return nil, customErr.NewRepositoryError("find", "extension", code, err)
	}
	return &ext, nil
}
