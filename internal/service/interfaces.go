package service

import (
	"context"

	"github.com/R4yL-dev/pkmc/internal/models"
)

type ItemService interface {
	CreateItem(ctx context.Context, extCode, langCode, typeName string, price *float64) (*models.Item, error)
}
