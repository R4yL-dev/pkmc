package service

import (
	"fmt"

	"github.com/R4yL-dev/pkmc/internal/models"
	"github.com/R4yL-dev/pkmc/internal/repository"
	"gorm.io/gorm"
)

type ItemService struct {
	db *gorm.DB
}

func NewItemService(db *gorm.DB) *ItemService {
	return &ItemService{db: db}
}

func (s *ItemService) CreateItem(extCode, langCode, typeName string, price *float64) (*models.Item, error) {
	var createdItem *models.Item

	err := repository.WithTransaction(s.db, func(tx *gorm.DB) error {
		extRepo := repository.NewExtensionRepository(tx)
		langRepo := repository.NewLanguageRepository(tx)
		typeRepo := repository.NewItemTypeRepository(tx)
		itemRepo := repository.NewItemRepository(tx)

		ext, err := extRepo.FindByCode(extCode)
		if err != nil {
			return fmt.Errorf("extension '%s' not found: %w", extCode, err)
		}

		lang, err := langRepo.FindByCode(langCode)
		if err != nil {
			return fmt.Errorf("language '%s' not found: %w", langCode, err)
		}

		itemType, err := typeRepo.FindByName(typeName)
		if err != nil {
			return fmt.Errorf("item type '%s' not found: %w", typeName, err)
		}

		item := &models.Item{
			ExtensionID: ext.ID,
			TypeID:      itemType.ID,
			LanguageID:  lang.ID,
			Price:       price,
		}

		if err := itemRepo.Create(item); err != nil {
			return fmt.Errorf("failed to create item: %w", err)
		}

		createdItem, err = itemRepo.FindByID(item.ID)
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
