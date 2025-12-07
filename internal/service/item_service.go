package service

import (
	"fmt"

	"github.com/R4yL-dev/pkmc/internal/models"
	"github.com/R4yL-dev/pkmc/internal/repository"
)

type itemService struct {
	uow repository.UnitOfWork
}

func NewItemService(uow repository.UnitOfWork) ItemService {
	return &itemService{uow: uow}
}

func (s *itemService) CreateItem(extCode, langCode, typeName string, price *float64) (*models.Item, error) {
	var createdItem *models.Item

	err := s.uow.Do(func(uow repository.UnitOfWork) error {
		ext, err := uow.Extensions().FindByCode(extCode)
		if err != nil {
			return fmt.Errorf("extension '%s' not found: %w", extCode, err)
		}

		lang, err := uow.Languages().FindByCode(langCode)
		if err != nil {
			return fmt.Errorf("language '%s' not found: %w", langCode, err)
		}

		itemType, err := uow.ItemTypes().FindByName(typeName)
		if err != nil {
			return fmt.Errorf("item type '%s' not found: %w", typeName, err)
		}

		item := &models.Item{
			ExtensionID: ext.ID,
			TypeID:      itemType.ID,
			LanguageID:  lang.ID,
			Price:       price,
		}

		if err := uow.Items().Create(item); err != nil {
			return fmt.Errorf("failed to create item: %w", err)
		}

		createdItem, err = uow.Items().FindByID(item.ID)
		if err != nil {
			return fmt.Errorf("failed to load created item: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdItem, nil
}
