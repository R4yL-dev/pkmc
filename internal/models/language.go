package models

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Code  string `gorm:"type:varchar(10);uniqueIndex;not null"`
	Name  string `gorm:"type:varchar(100);not null"`
	Items []Item `gorm:"foreignKey:LanguageID"`
}
