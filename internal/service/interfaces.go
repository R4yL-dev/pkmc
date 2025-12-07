package service

import "github.com/R4yL-dev/pkmc/internal/models"

type ItemService interface {
	CreateItem(extCode, langCode, typeName string, price *float64) (*models.Item, error)
}
