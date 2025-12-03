package models

import (
	"time"

	"gorm.io/gorm"
)

type Extension struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Code        string `gorm:"type:varchar(50);uniqueIndex"`
	BlockID     uint   `gorm:"not null;index"`
	Block       Block  `gorm:"foreignKey:BlockID;constraint:OnDelete:RESTRICT"`
	ReleaseDate *time.Time
	Items       []Item `gorm:"foreignKey:ExtensionID"`
}
