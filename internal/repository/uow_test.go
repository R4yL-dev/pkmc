package repository

import (
	"errors"
	"testing"

	"github.com/R4yL-dev/pkmc/internal/models"
	"github.com/R4yL-dev/pkmc/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitOfWork_DoCommit(t *testing.T) {
	// Setup
	db := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(t, db)

	uow := NewUnitOfWork(db)

	// Execute - successful transaction
	var createdItemID uint
	err := uow.Do(func(uow UnitOfWork) error {
		item := &models.Item{
			ExtensionID: 1,
			TypeID:      1,
			LanguageID:  1,
			Price:       testutil.FloatPtr(99.99),
		}

		err := uow.Items().Create(item)
		if err != nil {
			return err
		}

		createdItemID = item.ID
		return nil
	})

	// Assert
	assert.NoError(t, err)
	assert.NotZero(t, createdItemID)

	// Verify item was committed to database
	var item models.Item
	err = db.First(&item, createdItemID).Error
	assert.NoError(t, err)
	assert.Equal(t, createdItemID, item.ID)
}

func TestUnitOfWork_DoRollback(t *testing.T) {
	// Setup
	db := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(t, db)

	uow := NewUnitOfWork(db)

	// Execute - failing transaction
	var attemptedItemID uint
	err := uow.Do(func(uow UnitOfWork) error {
		item := &models.Item{
			ExtensionID: 1,
			TypeID:      1,
			LanguageID:  1,
			Price:       testutil.FloatPtr(99.99),
		}

		err := uow.Items().Create(item)
		if err != nil {
			return err
		}

		attemptedItemID = item.ID

		// Force an error to trigger rollback
		return errors.New("simulated error")
	})

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "simulated error")
	assert.NotZero(t, attemptedItemID)

	// Verify item was NOT committed to database
	var item models.Item
	err = db.First(&item, attemptedItemID).Error
	assert.Error(t, err, "Item should not exist after rollback")
}

func TestUnitOfWork_MultipleOperations(t *testing.T) {
	// Setup
	db := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(t, db)

	uow := NewUnitOfWork(db)

	// Execute - multiple operations in single transaction
	var item1ID, item2ID uint
	err := uow.Do(func(uow UnitOfWork) error {
		// Create first item
		item1 := &models.Item{
			ExtensionID: 1,
			TypeID:      1,
			LanguageID:  1,
			Price:       testutil.FloatPtr(99.99),
		}
		if err := uow.Items().Create(item1); err != nil {
			return err
		}
		item1ID = item1.ID

		// Create second item
		item2 := &models.Item{
			ExtensionID: 2,
			TypeID:      2,
			LanguageID:  2,
			Price:       testutil.FloatPtr(149.99),
		}
		if err := uow.Items().Create(item2); err != nil {
			return err
		}
		item2ID = item2.ID

		return nil
	})

	// Assert
	assert.NoError(t, err)
	assert.NotZero(t, item1ID)
	assert.NotZero(t, item2ID)

	// Verify both items were committed
	var item1, item2 models.Item
	err = db.First(&item1, item1ID).Error
	assert.NoError(t, err)
	err = db.First(&item2, item2ID).Error
	assert.NoError(t, err)
}

func TestUnitOfWork_RollbackMultipleOperations(t *testing.T) {
	// Setup
	db := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(t, db)

	uow := NewUnitOfWork(db)

	// Execute - multiple operations with failure
	var item1ID, item2ID uint
	err := uow.Do(func(uow UnitOfWork) error {
		// Create first item (should succeed)
		item1 := &models.Item{
			ExtensionID: 1,
			TypeID:      1,
			LanguageID:  1,
			Price:       testutil.FloatPtr(99.99),
		}
		if err := uow.Items().Create(item1); err != nil {
			return err
		}
		item1ID = item1.ID

		// Create second item (should succeed)
		item2 := &models.Item{
			ExtensionID: 2,
			TypeID:      2,
			LanguageID:  2,
			Price:       testutil.FloatPtr(149.99),
		}
		if err := uow.Items().Create(item2); err != nil {
			return err
		}
		item2ID = item2.ID

		// Force error to rollback both items
		return errors.New("rollback test")
	})

	// Assert
	assert.Error(t, err)
	assert.NotZero(t, item1ID)
	assert.NotZero(t, item2ID)

	// Verify neither item was committed
	var count int64
	db.Model(&models.Item{}).Count(&count)
	assert.Zero(t, count, "No items should exist after rollback")
}

func TestUnitOfWork_RepositoryAccess(t *testing.T) {
	// Setup
	db := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(t, db)

	uow := NewUnitOfWork(db)

	// Test all repository accessors
	err := uow.Do(func(uow UnitOfWork) error {
		// Verify all repositories are accessible
		assert.NotNil(t, uow.Items())
		assert.NotNil(t, uow.Extensions())
		assert.NotNil(t, uow.Languages())
		assert.NotNil(t, uow.ItemTypes())

		// Test actual usage
		ext, err := uow.Extensions().FindByCode("DRI")
		require.NoError(t, err)
		assert.NotNil(t, ext)

		lang, err := uow.Languages().FindByCode("fr")
		require.NoError(t, err)
		assert.NotNil(t, lang)

		itemType, err := uow.ItemTypes().FindByName("Display")
		require.NoError(t, err)
		assert.NotNil(t, itemType)

		return nil
	})

	assert.NoError(t, err)
}

func TestUnitOfWork_NestedTransactionBehavior(t *testing.T) {
	// Setup
	db := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(t, db)

	uow := NewUnitOfWork(db)

	// Test that inner error propagates correctly
	err := uow.Do(func(uow UnitOfWork) error {
		// First operation succeeds
		item1 := &models.Item{
			ExtensionID: 1,
			TypeID:      1,
			LanguageID:  1,
			Price:       testutil.FloatPtr(99.99),
		}
		if err := uow.Items().Create(item1); err != nil {
			return err
		}

		// Second operation fails with FK constraint
		item2 := &models.Item{
			ExtensionID: 9999, // Invalid FK
			TypeID:      1,
			LanguageID:  1,
			Price:       testutil.FloatPtr(99.99),
		}
		if err := uow.Items().Create(item2); err != nil {
			return err // Should trigger rollback
		}

		return nil
	})

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "FOREIGN KEY constraint failed")

	// Verify first item was also rolled back
	var count int64
	db.Model(&models.Item{}).Count(&count)
	assert.Zero(t, count, "All operations should be rolled back")
}
