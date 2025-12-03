package seed

import (
	"time"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	if err := SeedItemTypes(db); err != nil {
		return err
	}
	if err := SeedBlocks(db); err != nil {
		return err
	}
	if err := SeedExtension(db); err != nil {
		return err
	}
	if err := SeedLanguages(db); err != nil {
		return err
	}
	return nil
}

func datePtr(y int, m time.Month, d int) *time.Time {
	v := time.Date(y, m, d, -1, 0, 0, 0, time.UTC)
	return &v
}
