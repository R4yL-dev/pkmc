package repository

import "gorm.io/gorm"

type UnitOfWork struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) *UnitOfWork {
	return &UnitOfWork{db: db}
}

func (u *UnitOfWork) Begin() error {
	u.tx = u.db.Begin()
	return u.tx.Error
}

func (u *UnitOfWork) Commit() error {
	if u.tx == nil {
		return nil
	}
	return u.tx.Commit().Error
}

func (u *UnitOfWork) Rollback() error {
	if u.tx == nil {
		return nil
	}
	return u.tx.Rollback().Error
}

func (u *UnitOfWork) Items() *ItemRepository {
	db := u.db

	if u.tx != nil {
		db = u.tx
	}
	return NewItemRepository(db)
}

func (u *UnitOfWork) Extensions() *ExtensionRepository {
	db := u.db

	if u.tx != nil {
		db = u.tx
	}
	return NewExtensionRepository(db)
}

func (u *UnitOfWork) Languages() *LanguageRepository {
	db := u.db

	if u.tx != nil {
		db = u.tx
	}
	return NewLanguageRepository(db)
}

func (u *UnitOfWork) ItemTypes() *ItemTypeRepository {
	db := u.db

	if u.tx != nil {
		db = u.tx
	}
	return NewItemTypeRepository(db)
}

func WithTransaction(db *gorm.DB, fn func(*gorm.DB) error) error {
	return db.Transaction(fn)
}
