package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ExtensionID uint      `gorm:"not null;index"`
	Extension   Extension `gorm:"foreignKey:ExtensionID;constraint:OnDelete:RESTRICT"`
	TypeID      uint      `gorm:"not null;index"`
	Type        ItemType  `gorm:"foreignKey:TypeID;constraint:OnDelete:RESTRICT"`
	LanguageID  uint      `gorm:"not null;index"`
	Language    Language  `gorm:"foreignKey:LanguageID;constraint:OnDelete:RESTRICT"`
	Price       *float64  `gorm:"type:decimal(10,2)"`
}
