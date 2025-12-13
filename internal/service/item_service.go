package service

import (
	"context"
	"fmt"

	customErr "github.com/R4yL-dev/pkmc/internal/errors"
	"github.com/R4yL-dev/pkmc/internal/models"
	"github.com/R4yL-dev/pkmc/internal/repository"
)

type itemService struct {
	uow repository.UnitOfWork
}

func NewItemService(uow repository.UnitOfWork) ItemService {
	return &itemService{uow: uow}
}

func (s *itemService) CreateItem(ctx context.Context, extCode, langCode, typeName string, price *float64) (*models.Item, error) {
	var createdItem *models.Item

	err := s.uow.Do(ctx, func(uow repository.UnitOfWork) error {
		ext, err := uow.Extensions().FindByCode(ctx, extCode)
		if err != nil {
			return customErr.NewServiceError("create_item", "item_service", fmt.Sprintf("extension '%s' not found", extCode), err)
		}

		lang, err := uow.Languages().FindByCode(ctx, langCode)
		if err != nil {
			return customErr.NewServiceError("create_item", "item_service", fmt.Sprintf("language '%s' not found", langCode), err)
		}

		itemType, err := uow.ItemTypes().FindByName(ctx, typeName)
		if err != nil {
			return customErr.NewServiceError("create_item", "item_service", fmt.Sprintf("item type '%s' not found", typeName), err)
		}

		item := &models.Item{
			ExtensionID: ext.ID,
			TypeID:      itemType.ID,
			LanguageID:  lang.ID,
			Price:       price,
		}

		if err := uow.Items().Create(ctx, item); err != nil {
			return customErr.NewServiceError("create_item", "item_service", "failed to create item", err)
		}

		createdItem, err = uow.Items().FindByID(ctx, item.ID)
		if err != nil {
			return customErr.NewServiceError("create_item", "item_service", "failed to load created item", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdItem, nil
}
