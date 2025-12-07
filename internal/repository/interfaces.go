package repository

import "github.com/R4yL-dev/pkmc/internal/models"

type ItemRepository interface {
	Create(item *models.Item) error
	FindByID(id uint) (*models.Item, error)
}

type ExtensionRepository interface {
	FindByCode(code string) (*models.Extension, error)
}

type LanguageRepository interface {
	FindByCode(code string) (*models.Language, error)
}

type ItemTypeRepository interface {
	FindByName(name string) (*models.ItemType, error)
}

type UnitOfWork interface {
	Do(fn func(uow UnitOfWork) error) error
	Items() ItemRepository
	Extensions() ExtensionRepository
	Languages() LanguageRepository
	ItemTypes() ItemTypeRepository
}
