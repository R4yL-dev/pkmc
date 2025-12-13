package repository

import (
	"context"
	"errors"

	customErr "github.com/R4yL-dev/pkmc/internal/errors"
	"github.com/R4yL-dev/pkmc/internal/models"
	"gorm.io/gorm"
)

type blockRepository struct {
	db *gorm.DB
}

func NewBlockRepository(db *gorm.DB) BlockRepository {
	return &blockRepository{db: db}
}

func (r *blockRepository) FindByCode(ctx context.Context, code string) (*models.Block, error) {
	var block models.Block

	err := r.db.WithContext(ctx).Where("code = ?", code).First(&block).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErr.NewRepositoryError("find", "block", code, customErr.ErrEntityNotFound)
		}
		return nil, customErr.NewRepositoryError("find", "block", code, err)
	}
	return &block, nil
}
