package repository

import (
	"context"
	"testing"
	"time"

	"github.com/R4yL-dev/pkmc/internal/models"
	"github.com/R4yL-dev/pkmc/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestItemRepository_Create(t *testing.T) {
	tests := []struct {
		name          string
		item          *models.Item
		expectedError bool
		errorContains string
	}{
		{
			name: "success - create item with price",
			item: &models.Item{
				ExtensionID: 1, // From seed data (DRI)
				TypeID:      1, // From seed data (ETB)
				LanguageID:  1, // From seed data (fr)
				Price:       testutil.FloatPtr(129.99),
			},
			expectedError: false,
		},
		{
			name: "success - create item without price",
			item: &models.Item{
				ExtensionID: 1,
				TypeID:      1,
				LanguageID:  1,
				Price:       nil,
			},
			expectedError: false,
		},
		{
			name: "error - invalid extension id (FK constraint)",
			item: &models.Item{
				ExtensionID: 9999,
				TypeID:      1,
				LanguageID:  1,
				Price:       testutil.FloatPtr(99.99),
			},
			expectedError: true,
			errorContains: "FOREIGN KEY constraint failed",
		},
		{
			name: "error - invalid type id (FK constraint)",
			item: &models.Item{
				ExtensionID: 1,
				TypeID:      9999,
				LanguageID:  1,
				Price:       testutil.FloatPtr(99.99),
			},
			expectedError: true,
			errorContains: "FOREIGN KEY constraint failed",
		},
		{
			name: "error - invalid language id (FK constraint)",
			item: &models.Item{
				ExtensionID: 1,
				TypeID:      1,
				LanguageID:  9999,
				Price:       testutil.FloatPtr(99.99),
			},
			expectedError: true,
			errorContains: "FOREIGN KEY constraint failed",
		},
		{
			name: "error - missing extension id",
			item: &models.Item{
				ExtensionID: 0,
				TypeID:      1,
				LanguageID:  1,
				Price:       testutil.FloatPtr(99.99),
			},
			expectedError: true,
			errorContains: "FOREIGN KEY constraint failed",
		},
		{
			name: "error - missing type id",
			item: &models.Item{
				ExtensionID: 1,
				TypeID:      0,
				LanguageID:  1,
				Price:       testutil.FloatPtr(99.99),
			},
			expectedError: true,
			errorContains: "FOREIGN KEY constraint failed",
		},
		{
			name: "error - missing language id",
			item: &models.Item{
				ExtensionID: 1,
				TypeID:      1,
				LanguageID:  0,
				Price:       testutil.FloatPtr(99.99),
			},
			expectedError: true,
			errorContains: "FOREIGN KEY constraint failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			db := testutil.SetupTestDB(t)
			defer testutil.CleanupTestDB(t, db)

			repo := NewItemRepository(db)
			ctx := context.Background()

			// Execute
			err := repo.Create(ctx, tt.item)

			// Assert
			if tt.expectedError {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
				assert.NotZero(t, tt.item.ID, "ID should be assigned after creation")
				assert.NotZero(t, tt.item.CreatedAt)
				assert.NotZero(t, tt.item.UpdatedAt)
			}
		})
	}
}

func TestItemRepository_Create_WithTimeout(t *testing.T) {
	// Setup
	db := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(t, db)

	repo := NewItemRepository(db)

	// Create a context with very short timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()

	// Wait to ensure timeout
	time.Sleep(2 * time.Millisecond)

	item := &models.Item{
		ExtensionID: 1,
		TypeID:      1,
		LanguageID:  1,
		Price:       testutil.FloatPtr(99.99),
	}

	// Execute
	err := repo.Create(ctx, item)

	// Assert - should fail with context deadline exceeded
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
}

func TestItemRepository_FindByID(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*gorm.DB) uint
		itemID        uint
		expectedError bool
		validate      func(*testing.T, *models.Item)
	}{
		{
			name: "success - find existing item with associations",
			setup: func(db *gorm.DB) uint {
				item := &models.Item{
					ExtensionID: 1,
					TypeID:      1,
					LanguageID:  1,
					Price:       testutil.FloatPtr(149.99),
				}
				db.Create(item)
				return item.ID
			},
			expectedError: false,
			validate: func(t *testing.T, item *models.Item) {
				assert.NotNil(t, item)
				assert.NotZero(t, item.ID)
				assert.Equal(t, uint(1), item.ExtensionID)
				assert.Equal(t, uint(1), item.TypeID)
				assert.Equal(t, uint(1), item.LanguageID)
				assert.NotNil(t, item.Price)
				assert.InDelta(t, 149.99, *item.Price, 0.001)

				// Verify associations are preloaded
				assert.NotZero(t, item.Extension.ID, "Extension should be preloaded")
				assert.NotEmpty(t, item.Extension.Code, "Extension should be preloaded")
				assert.NotZero(t, item.Type.ID, "Type should be preloaded")
				assert.NotEmpty(t, item.Type.Name, "Type should be preloaded")
				assert.NotZero(t, item.Language.ID, "Language should be preloaded")
				assert.NotEmpty(t, item.Language.Code, "Language should be preloaded")
			},
		},
		{
			name: "success - find item without price",
			setup: func(db *gorm.DB) uint {
				item := &models.Item{
					ExtensionID: 1,
					TypeID:      1,
					LanguageID:  1,
					Price:       nil,
				}
				db.Create(item)
				return item.ID
			},
			expectedError: false,
			validate: func(t *testing.T, item *models.Item) {
				assert.NotNil(t, item)
				assert.Nil(t, item.Price)
			},
		},
		{
			name: "error - item not found",
			setup: func(db *gorm.DB) uint {
				return 9999 // Non-existent ID
			},
			expectedError: true,
			validate: func(t *testing.T, item *models.Item) {
				assert.Nil(t, item)
			},
		},
		{
			name: "error - zero id",
			setup: func(db *gorm.DB) uint {
				return 0
			},
			expectedError: true,
			validate: func(t *testing.T, item *models.Item) {
				assert.Nil(t, item)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			db := testutil.SetupTestDB(t)
			defer testutil.CleanupTestDB(t, db)

			repo := NewItemRepository(db)
			ctx := context.Background()

			var itemID uint
			if tt.setup != nil {
				itemID = tt.setup(db)
			} else {
				itemID = tt.itemID
			}

			// Execute
			item, err := repo.FindByID(ctx, itemID)

			// Assert
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.validate != nil {
				tt.validate(t, item)
			}
		})
	}
}

func TestItemRepository_FindByID_WithTimeout(t *testing.T) {
	// Setup
	db := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(t, db)

	repo := NewItemRepository(db)

	// Create an item first
	item := &models.Item{
		ExtensionID: 1,
		TypeID:      1,
		LanguageID:  1,
		Price:       testutil.FloatPtr(99.99),
	}
	db.Create(item)

	// Create a context with very short timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()

	// Wait to ensure timeout
	time.Sleep(2 * time.Millisecond)

	// Execute
	_, err := repo.FindByID(ctx, item.ID)

	// Assert - should fail with context deadline exceeded
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
}

func TestItemRepository_MultipleItems(t *testing.T) {
	// Setup
	db := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(t, db)

	repo := NewItemRepository(db)
	ctx := context.Background()

	// Create multiple items
	items := []*models.Item{
		{ExtensionID: 1, TypeID: 1, LanguageID: 1, Price: testutil.FloatPtr(99.99)},
		{ExtensionID: 1, TypeID: 2, LanguageID: 1, Price: testutil.FloatPtr(149.99)},
		{ExtensionID: 2, TypeID: 1, LanguageID: 2, Price: nil},
	}

	for i, item := range items {
		err := repo.Create(ctx, item)
		require.NoError(t, err, "Failed to create item %d", i)
		assert.NotZero(t, item.ID)
	}

	// Verify all items can be retrieved
	for i, item := range items {
		retrieved, err := repo.FindByID(ctx, item.ID)
		require.NoError(t, err, "Failed to find item %d", i)
		assert.Equal(t, item.ID, retrieved.ID)
		testutil.AssertItemEqual(t, item, retrieved)
	}
}
