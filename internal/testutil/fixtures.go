package testutil

import (
	"github.com/R4yL-dev/pkmc/internal/models"
)

func CreateTestBlock(overrides ...func(*models.Block)) *models.Block {
	block := &models.Block{
		Name:        "Test Block",
		Code:        "TST",
		ReleaseDate: DatePtr(2024, 1, 1),
	}

	for _, override := range overrides {
		override(block)
	}

	return block
}

func CreateTestExtension(blockID uint, overrides ...func(*models.Extension)) *models.Extension {
	ext := &models.Extension{
		Name:        "Test Extension",
		Code:        "TEX",
		BlockID:     blockID,
		ReleaseDate: DatePtr(2024, 1, 1),
	}

	for _, override := range overrides {
		override(ext)
	}

	return ext
}

func CreateTestItemType(overrides ...func(*models.ItemType)) *models.ItemType {
	itemType := &models.ItemType{
		Name: "Test Type",
	}

	for _, override := range overrides {
		override(itemType)
	}

	return itemType
}

func CreateTestLanguage(overrides ...func(*models.Language)) *models.Language {
	lang := &models.Language{
		Code: "tt",
		Name: "Test Language",
	}

	for _, override := range overrides {
		override(lang)
	}

	return lang
}

func CreateTestItem(extID, typeID, langID uint, overrides ...func(*models.Item)) *models.Item {
	item := &models.Item{
		ExtensionID: extID,
		TypeID:      typeID,
		LanguageID:  langID,
		Price:       FloatPtr(99.99),
	}

	for _, override := range overrides {
		override(item)
	}

	return item
}
