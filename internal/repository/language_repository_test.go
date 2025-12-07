package repository

import (
	"testing"

	"github.com/R4yL-dev/pkmc/internal/models"
	"github.com/R4yL-dev/pkmc/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestLanguageRepository_FindByCode(t *testing.T) {
	tests := []struct {
		name          string
		code          string
		expectedError bool
		validate      func(*testing.T, *models.Language)
	}{
		{
			name:          "success - find French",
			code:          "fr",
			expectedError: false,
			validate: func(t *testing.T, lang *models.Language) {
				assert.NotNil(t, lang)
				assert.Equal(t, "fr", lang.Code)
				assert.Equal(t, "Français", lang.Name)
				assert.NotZero(t, lang.ID)
			},
		},
		{
			name:          "success - find English",
			code:          "en",
			expectedError: false,
			validate: func(t *testing.T, lang *models.Language) {
				assert.NotNil(t, lang)
				assert.Equal(t, "en", lang.Code)
				assert.Equal(t, "English", lang.Name)
			},
		},
		{
			name:          "success - find German",
			code:          "de",
			expectedError: false,
			validate: func(t *testing.T, lang *models.Language) {
				assert.NotNil(t, lang)
				assert.Equal(t, "de", lang.Code)
				assert.Equal(t, "Deutsch", lang.Name)
			},
		},
		{
			name:          "success - find Spanish",
			code:          "es",
			expectedError: false,
			validate: func(t *testing.T, lang *models.Language) {
				assert.NotNil(t, lang)
				assert.Equal(t, "es", lang.Code)
				assert.Equal(t, "Español", lang.Name)
			},
		},
		{
			name:          "error - language not found",
			code:          "invalid",
			expectedError: true,
			validate: func(t *testing.T, lang *models.Language) {
				assert.Nil(t, lang)
			},
		},
		{
			name:          "error - empty code",
			code:          "",
			expectedError: true,
			validate: func(t *testing.T, lang *models.Language) {
				assert.Nil(t, lang)
			},
		},
		{
			name:          "error - case sensitivity",
			code:          "FR", // uppercase
			expectedError: true,
			validate: func(t *testing.T, lang *models.Language) {
				assert.Nil(t, lang)
			},
		},
		{
			name:          "error - code with spaces",
			code:          "fr ",
			expectedError: true,
			validate: func(t *testing.T, lang *models.Language) {
				assert.Nil(t, lang)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			db := testutil.SetupTestDB(t)
			defer testutil.CleanupTestDB(t, db)

			repo := NewLanguageRepository(db)

			// Execute
			lang, err := repo.FindByCode(tt.code)

			// Assert
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.validate != nil {
				tt.validate(t, lang)
			}
		})
	}
}

func TestLanguageRepository_CustomLanguage(t *testing.T) {
	// Setup
	db := testutil.SetupTestDB(t)
	defer testutil.CleanupTestDB(t, db)

	// Create a custom language
	customLang := testutil.CreateTestLanguage(func(l *models.Language) {
		l.Code = "it"
		l.Name = "Italiano"
	})

	err := db.Create(customLang).Error
	assert.NoError(t, err)

	repo := NewLanguageRepository(db)

	// Find custom language
	found, err := repo.FindByCode("it")
	assert.NoError(t, err)
	assert.NotNil(t, found)
	testutil.AssertLanguageEqual(t, customLang, found)
}
