package models

import (
	"time"

	"gorm.io/gorm"
)

type Block struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);uniqueIndex;not null"`
	ReleaseDate *time.Time
	Code        string      `gorm:"type:varchar(50);uniqueIndex;not null"`
	Extensions  []Extension `gorm:"foreignKey:BlockID"`
}
