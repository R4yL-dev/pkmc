package repository

import "gorm.io/gorm"

type unitOfWork struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	return &unitOfWork{db: db}
}

func (u *unitOfWork) Do(fn func(uow UnitOfWork) error) error {
	tx := u.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	txUoW := &unitOfWork{
		db: u.db,
		tx: tx,
	}

	if err := fn(txUoW); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (u *unitOfWork) Items() ItemRepository {
	db := u.db

	if u.tx != nil {
		db = u.tx
	}
	return NewItemRepository(db)
}

func (u *unitOfWork) Extensions() ExtensionRepository {
	db := u.db

	if u.tx != nil {
		db = u.tx
	}
	return NewExtensionRepository(db)
}

func (u *unitOfWork) Languages() LanguageRepository {
	db := u.db

	if u.tx != nil {
		db = u.tx
	}
	return NewLanguageRepository(db)
}

func (u *unitOfWork) ItemTypes() ItemTypeRepository {
	db := u.db

	if u.tx != nil {
		db = u.tx
	}
	return NewItemTypeRepository(db)
}
