package testutil

import (
	"testing"

	"github.com/R4yL-dev/pkmc/internal/models"
	"github.com/R4yL-dev/pkmc/internal/seed"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(t, err, "Failed to open test database")

	sqlDB, err := db.DB()
	require.NoError(t, err, "Failed to get underlying database")
	_, err = sqlDB.Exec("PRAGMA foreign_keys = ON")
	require.NoError(t, err, "Failed to enable foreign keys")

	err = db.AutoMigrate(models.GetModels()...)
	require.NoError(t, err, "Failed to migrate test database")

	seed.Seed(db)

	return db
}

func SetupTestDBWithoutSeed(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(t, err, "Failed to open test database")

	sqlDB, err := db.DB()
	require.NoError(t, err, "Failed to get underlying database")
	_, err = sqlDB.Exec("PRAGMA foreign_keys = ON")
	require.NoError(t, err, "Failed to enable foreign keys")

	err = db.AutoMigrate(models.GetModels()...)
	require.NoError(t, err, "Failed to migrate test database")

	return db
}

func CleanupTestDB(t *testing.T, db *gorm.DB) {
	t.Helper()

	sqlDB, err := db.DB()
	require.NoError(t, err, "Failed to get underlying database")

	err = sqlDB.Close()
	require.NoError(t, err, "Failed to close test database")
}
