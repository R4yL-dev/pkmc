package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	customErr "github.com/R4yL-dev/pkmc/internal/errors"
)

func InitDB(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, customErr.NewDBError("open", err, dbPath)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, customErr.NewDBError("get_sql_db", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, customErr.NewDBError("ping", err)
	}

	return db, nil
}

func CloseDB(db *gorm.DB) error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return customErr.NewDBError("get_sql_db", err)
	}

	if err := sqlDB.Close(); err != nil {
		return customErr.NewDBError("close", err)
	}
	return nil
}
