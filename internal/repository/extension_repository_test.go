package repository

import (
	"testing"

	"github.com/R4yL-dev/pkmc/internal/models"
	"github.com/R4yL-dev/pkmc/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestExtensionRepository_FindByCode(t *testing.T) {
	tests := []struct {
		name          string
		code          string
		expectedError bool
		validate      func(*testing.T, *models.Extension)
	}{
		{
			name:          "success - find existing extension (from seed)",
			code:          "DRI",
			expectedError: false,
			validate: func(t *testing.T, ext *models.Extension) {
				assert.NotNil(t, ext)
				assert.Equal(t, "DRI", ext.Code)
				assert.Equal(t, "Rivalités Destinées", ext.Name)
				assert.NotZero(t, ext.ID)
				assert.NotZero(t, ext.BlockID)
			},
		},
		{
			name:          "success - find another extension",
			code:          "EVS",
			expectedError: false,
			validate: func(t *testing.T, ext *models.Extension) {
				assert.NotNil(t, ext)
				assert.Equal(t, "EVS", ext.Code)
				assert.Equal(t, "Évolution Céleste", ext.Name)
				assert.NotZero(t, ext.ID)
			},
		},
		{
			name:          "error - extension not found",
			code:          "INVALID",
			expectedError: true,
			validate: func(t *testing.T, ext *models.Extension) {
				assert.Nil(t, ext)
			},
		},
		{
			name:          "error - empty code",
			code:          "",
			expectedError: true,
			validate: func(t *testing.T, ext *models.Extension) {
				assert.Nil(t, ext)
			},
		},
		{
			name:          "error - case sensitivity",
			code:          "dri", // lowercase
			expectedError: true,
			validate: func(t *testing.T, ext *models.Extension) {
				assert.Nil(t, ext)
			},
		},
		{
			name:          "error - code with spaces",
			code:          "DRI ",
			expectedError: true,
			validate: func(t *testing.T, ext *models.Extension) {
				assert.Nil(t, ext)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			db := testutil.SetupTestDB(t)
			defer testutil.CleanupTestDB(t, db)

			repo := NewExtensionRepository(db)

			// Execute
			ext, err := repo.FindByCode(tt.code)

			// Assert
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.validate != nil {
				tt.validate(t, ext)
			}
		})
	}
}

func TestExtensionRepository_CustomExtension(t *testing.T) {
	// Setup
	db := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(t, db)

	// Create a custom extension
	customExt := testutil.CreateTestExtension(1, func(e *models.Extension) {
		e.Code = "CUSTOM"
		e.Name = "Custom Extension"
	})

	err := db.Create(customExt).Error
	assert.NoError(t, err)

	repo := NewExtensionRepository(db)

	// Find custom extension
	found, err := repo.FindByCode("CUSTOM")
	assert.NoError(t, err)
	assert.NotNil(t, found)
	testutil.AssertExtensionEqual(t, customExt, found)
}
