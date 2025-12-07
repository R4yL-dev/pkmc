package repository

import (
	"testing"

	"github.com/R4yL-dev/pkmc/internal/models"
	"github.com/R4yL-dev/pkmc/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestItemTypeRepository_FindByName(t *testing.T) {
	tests := []struct {
		name          string
		typeName      string
		expectedError bool
		validate      func(*testing.T, *models.ItemType)
	}{
		{
			name:          "success - find ETB",
			typeName:      "ETB",
			expectedError: false,
			validate: func(t *testing.T, itemType *models.ItemType) {
				assert.NotNil(t, itemType)
				assert.Equal(t, "ETB", itemType.Name)
				assert.NotZero(t, itemType.ID)
			},
		},
		{
			name:          "success - find Display",
			typeName:      "Display",
			expectedError: false,
			validate: func(t *testing.T, itemType *models.ItemType) {
				assert.NotNil(t, itemType)
				assert.Equal(t, "Display", itemType.Name)
				assert.NotZero(t, itemType.ID)
			},
		},
		{
			name:          "success - find Bundle",
			typeName:      "Bundle",
			expectedError: false,
			validate: func(t *testing.T, itemType *models.ItemType) {
				assert.NotNil(t, itemType)
				assert.Equal(t, "Bundle", itemType.Name)
			},
		},
		{
			name:          "success - find Booster",
			typeName:      "Booster",
			expectedError: false,
			validate: func(t *testing.T, itemType *models.ItemType) {
				assert.NotNil(t, itemType)
				assert.Equal(t, "Booster", itemType.Name)
			},
		},
		{
			name:          "success - find Sleeve Booster",
			typeName:      "Sleeve Booster",
			expectedError: false,
			validate: func(t *testing.T, itemType *models.ItemType) {
				assert.NotNil(t, itemType)
				assert.Equal(t, "Sleeve Booster", itemType.Name)
			},
		},
		{
			name:          "error - type not found",
			typeName:      "InvalidType",
			expectedError: true,
			validate: func(t *testing.T, itemType *models.ItemType) {
				assert.Nil(t, itemType)
			},
		},
		{
			name:          "error - empty name",
			typeName:      "",
			expectedError: true,
			validate: func(t *testing.T, itemType *models.ItemType) {
				assert.Nil(t, itemType)
			},
		},
		{
			name:          "error - case sensitivity",
			typeName:      "etb", // lowercase
			expectedError: true,
			validate: func(t *testing.T, itemType *models.ItemType) {
				assert.Nil(t, itemType)
			},
		},
		{
			name:          "error - name with extra spaces",
			typeName:      "ETB ",
			expectedError: true,
			validate: func(t *testing.T, itemType *models.ItemType) {
				assert.Nil(t, itemType)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			db := testutil.SetupTestDB(t)
			defer testutil.CleanupTestDB(t, db)

			repo := NewItemTypeRepository(db)

			// Execute
			itemType, err := repo.FindByName(tt.typeName)

			// Assert
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.validate != nil {
				tt.validate(t, itemType)
			}
		})
	}
}

func TestItemTypeRepository_CustomItemType(t *testing.T) {
	// Setup
	db := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(t, db)

	// Create a custom item type
	customType := testutil.CreateTestItemType(func(it *models.ItemType) {
		it.Name = "Custom Type"
	})

	err := db.Create(customType).Error
	assert.NoError(t, err)

	repo := NewItemTypeRepository(db)

	// Find custom item type
	found, err := repo.FindByName("Custom Type")
	assert.NoError(t, err)
	assert.NotNil(t, found)
	testutil.AssertItemTypeEqual(t, customType, found)
}
