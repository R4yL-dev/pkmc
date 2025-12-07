package repository

import (
	"github.com/R4yL-dev/pkmc/internal/models"
	"gorm.io/gorm"
)

type languageRepository struct {
	db *gorm.DB
}

func NewLanguageRepository(db *gorm.DB) LanguageRepository {
	return &languageRepository{db: db}
}

func (r *languageRepository) FindByCode(code string) (*models.Language, error) {
	var lang models.Language

	err := r.db.Where("code = ?", code).First(&lang).Error
	if err != nil {
		return nil, err
	}
	return &lang, nil
}
