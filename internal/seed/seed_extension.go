package seed

import (
	"log"
	"time"

	"github.com/R4yL-dev/pkmc/internal/models"
	"gorm.io/gorm"
)

func SeedExtension(db *gorm.DB) error {
	var blocks []models.Block
	if err := db.Find(&blocks).Error; err != nil {
		return err
	}

	blockMap := make(map[string]models.Block)
	for _, b := range blocks {
		blockMap[b.Code] = b
	}

	extensions := []models.Extension{
		// EB
		{Name: "Épée et Bouclier", Code: "SSH", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2020, time.February, 7)},
		{Name: "Promos Épée et Bouclier", Code: "SWSH", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2020, time.February, 7)},
		{Name: "Clash des Rebelles", Code: "RCL", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2020, time.May, 1)},
		{Name: "Ténèbres Embrasées", Code: "DAA", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2020, time.August, 14)},
		{Name: "La Voie du Maître", Code: "CPA", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2020, time.September, 25)},
		{Name: "Voltage Éclatant", Code: "VIV", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2020, time.November, 13)},
		{Name: "Destinées Radieuses", Code: "SHF", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2021, time.February, 19)},
		{Name: "Styles de Combat", Code: "BST", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2021, time.March, 19)},
		{Name: "Règne de Glace", Code: "CRE", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2021, time.June, 18)},
		{Name: "Évolution Céleste", Code: "EVS", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2021, time.August, 27)},
		{Name: "Célébrations", Code: "CEL", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2021, time.October, 8)},
		{Name: "Poing de Fusion", Code: "FST", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2021, time.November, 12)},
		{Name: "Stars Étincelantes", Code: "BRS", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2022, time.February, 25)},
		{Name: "Astres Radieux", Code: "ASR", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2022, time.May, 27)},
		{Name: "Pokémon GO", Code: "PGO", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2022, time.July, 1)},
		{Name: "Tempête Argentée", Code: "SIT", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2022, time.November, 11)},
		{Name: "Zénith Suprême", Code: "CRZ", BlockID: blockMap["EB"].ID, ReleaseDate: datePtr(2023, time.February, 27)},

		// EV
		{Name: "Écarlate et Violet", Code: "SVI", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2023, time.March, 31)},
		{Name: "Promos Écarlate et Violet", Code: "SVP", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2023, time.March, 31)},
		{Name: "Évolutions à Paldea", Code: "PAL", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2023, time.June, 9)},
		{Name: "Flammes Obsidiennes", Code: "OBF", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2023, time.August, 11)},
		{Name: "151", Code: "MEW", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2023, time.September, 22)},
		{Name: "Faille Paradoxe", Code: "PAR", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2023, time.November, 3)},
		{Name: "Destinées de Paldea", Code: "PAF", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2024, time.January, 26)},
		{Name: "Forces Temporelles", Code: "TEF", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2024, time.March, 22)},
		{Name: "Mascarade Crépusculaire", Code: "TWM", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2024, time.May, 24)},
		{Name: "Fable Nébuleuse", Code: "SFA", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2024, time.August, 8)},
		{Name: "Couronne Stellaire", Code: "SCR", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2024, time.September, 13)},
		{Name: "Étincelles Déferlantes", Code: "SSP", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2024, time.November, 11)},
		{Name: "Évolutions Prismatiques", Code: "PRE", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2025, time.January, 17)},
		{Name: "Aventures Ensemble", Code: "JTG", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2025, time.March, 28)},
		{Name: "Rivalités Destinées", Code: "DRI", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2025, time.May, 30)},
		{Name: "Flamme Blanche", Code: "WHT", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2025, time.July, 18)},
		{Name: "Foudre Noire", Code: "BLK", BlockID: blockMap["EV"].ID, ReleaseDate: datePtr(2025, time.July, 18)},

		// ME
		{Name: "Méga-Évolution", Code: "MEG", BlockID: blockMap["ME"].ID, ReleaseDate: datePtr(2025, time.October, 10)},
		{Name: "Promos Méga-Évolution", Code: "MEP", BlockID: blockMap["ME"].ID, ReleaseDate: datePtr(2025, time.October, 10)},
		{Name: "Flammes Fantasmagoriques", Code: "PFL", BlockID: blockMap["ME"].ID, ReleaseDate: datePtr(2025, time.November, 14)},
	}

	for _, e := range extensions {
		var existing models.Extension

		if err := db.Where("code = ?", e.Code).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&e).Error; err != nil {
					return err
				}
				log.Printf("Seed: created extension '%s'\n", e.Name)
			} else {
				return err
			}
		}
	}
	return nil
}
