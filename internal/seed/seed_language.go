package seed

import (
	"log"

	"github.com/R4yL-dev/pkmc/internal/models"
	"gorm.io/gorm"
)

func SeedLanguages(db *gorm.DB) error {
	languages := []models.Language{
		{Code: "fr", Name: "Français"},
		{Code: "en", Name: "English"},
		{Code: "de", Name: "Deutsch"},
		{Code: "es", Name: "Español"},
	}

	for _, l := range languages {
		var existing models.Language

		if err := db.Where("code = ?", l.Code).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&l).Error; err != nil {
					return err
				}
				log.Printf("Seed: created language '%s'\n", l.Name)
			} else {
				return err
			}
		}
	}
	return nil
}
