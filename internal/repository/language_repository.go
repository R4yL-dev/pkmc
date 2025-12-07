package repository

import (
	"context"

	"github.com/R4yL-dev/pkmc/internal/models"
	"gorm.io/gorm"
)

type languageRepository struct {
	db *gorm.DB
}

func NewLanguageRepository(db *gorm.DB) LanguageRepository {
	return &languageRepository{db: db}
}

func (r *languageRepository) FindByCode(ctx context.Context, code string) (*models.Language, error) {
	var lang models.Language

	err := r.db.WithContext(ctx).Where("code = ?", code).First(&lang).Error
	if err != nil {
		return nil, err
	}
	return &lang, nil
}
