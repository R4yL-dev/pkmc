package repository

import (
	"context"

	"github.com/R4yL-dev/pkmc/internal/models"
)

type ItemRepository interface {
	Create(ctx context.Context, item *models.Item) error
	FindByID(ctx context.Context, id uint) (*models.Item, error)
}

type ExtensionRepository interface {
	FindByCode(ctx context.Context, code string) (*models.Extension, error)
}

type LanguageRepository interface {
	FindByCode(ctx context.Context, code string) (*models.Language, error)
}

type ItemTypeRepository interface {
	FindByName(ctx context.Context, name string) (*models.ItemType, error)
}

type BlockRepository interface {
	FindByCode(ctx context.Context, code string) (*models.Block, error)
}

type UnitOfWork interface {
	Do(ctx context.Context, fn func(uow UnitOfWork) error) error
	Items() ItemRepository
	Extensions() ExtensionRepository
	Languages() LanguageRepository
	ItemTypes() ItemTypeRepository
	Blocks() BlockRepository
}
