package seed

import (
	"time"

	"github.com/R4yL-dev/pkmc/internal/models"
	"gorm.io/gorm"
)

func SeedBlocks(db *gorm.DB) error {
	blocks := []models.Block{
		{Name: "Épée et Bouclier", Code: "EB", ReleaseDate: datePtr(2020, time.February, 7)},
		{Name: "Écarlate et Violet", Code: "EV", ReleaseDate: datePtr(2023, time.March, 31)},
		{Name: "Méga-Évolution", Code: "ME", ReleaseDate: datePtr(2025, time.October, 10)},
	}

	for _, b := range blocks {
		var existing models.Block

		if err := db.Where("name = ?", b.Name).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&b).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}
	return nil
}
