package models

func GetModels() []interface{} {
	return []interface{}{
		&Block{},
		&Extension{},
		&ItemType{},
		&Item{},
		&Language{},
	}
}
