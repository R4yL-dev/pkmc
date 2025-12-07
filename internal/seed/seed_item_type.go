package seed

import (
	"github.com/R4yL-dev/pkmc/internal/models"
	"gorm.io/gorm"
)

func SeedItemTypes(db *gorm.DB) error {
	itemTypes := []models.ItemType{
		{Name: "ETB"},
		{Name: "Display"},
		{Name: "Bundle"},
		{Name: "Booster"},
		{Name: "Sleeve Booster"},
	}

	for _, t := range itemTypes {
		var existing models.ItemType

		if err := db.Where("name = ?", t.Name).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&t).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}
	return nil
}
