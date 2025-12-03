package models

import "gorm.io/gorm"

type ItemType struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Items []Item `gorm:"foreignKey:TypeID"`
}
