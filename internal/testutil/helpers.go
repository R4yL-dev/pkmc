package testutil

import (
	"testing"
	"time"

	"github.com/R4yL-dev/pkmc/internal/models"
	"github.com/stretchr/testify/assert"
)

func FloatPtr(f float64) *float64 {
	return &f
}

func DatePtr(year int, month int, day int) *time.Time {
	releaseDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return &releaseDate
}

func AssertItemEqual(t *testing.T, expected, actual *models.Item, msgAndArgs ...interface{}) {
	t.Helper()

	if expected == nil && actual == nil {
		return
	}

	assert.NotNil(t, actual, msgAndArgs...)
	if actual == nil {
		return
	}

	assert.Equal(t, expected.ID, actual.ID, msgAndArgs...)
	assert.Equal(t, expected.ExtensionID, actual.ExtensionID, msgAndArgs...)
	assert.Equal(t, expected.TypeID, actual.TypeID, msgAndArgs...)
	assert.Equal(t, expected.LanguageID, actual.LanguageID, msgAndArgs...)

	// Compare prices (handle nil cases)
	if expected.Price == nil {
		assert.Nil(t, actual.Price, msgAndArgs...)
	} else {
		assert.NotNil(t, actual.Price, msgAndArgs...)
		if actual.Price != nil {
			assert.InDelta(t, *expected.Price, *actual.Price, 0.001, msgAndArgs...)
		}
	}

	// Compare associations if loaded
	if expected.Extension.ID != 0 {
		AssertExtensionEqual(t, &expected.Extension, &actual.Extension, msgAndArgs...)
	}
	if expected.Type.ID != 0 {
		AssertItemTypeEqual(t, &expected.Type, &actual.Type, msgAndArgs...)
	}
	if expected.Language.ID != 0 {
		AssertLanguageEqual(t, &expected.Language, &actual.Language, msgAndArgs...)
	}
}

func AssertExtensionEqual(t *testing.T, expected, actual *models.Extension, msgAndArgs ...interface{}) {
	t.Helper()

	if expected == nil && actual == nil {
		return
	}

	assert.NotNil(t, actual, msgAndArgs...)
	if actual == nil {
		return
	}

	assert.Equal(t, expected.ID, actual.ID, msgAndArgs...)
	assert.Equal(t, expected.Name, actual.Name, msgAndArgs...)
	assert.Equal(t, expected.Code, actual.Code, msgAndArgs...)
	assert.Equal(t, expected.BlockID, actual.BlockID, msgAndArgs...)
}

func AssertItemTypeEqual(t *testing.T, expected, actual *models.ItemType, msgAndArgs ...interface{}) {
	t.Helper()

	if expected == nil && actual == nil {
		return
	}

	assert.NotNil(t, actual, msgAndArgs...)
	if actual == nil {
		return
	}

	assert.Equal(t, expected.ID, actual.ID, msgAndArgs...)
	assert.Equal(t, expected.Name, actual.Name, msgAndArgs...)
}

func AssertLanguageEqual(t *testing.T, expected, actual *models.Language, msgAndArgs ...interface{}) {
	t.Helper()

	if expected == nil && actual == nil {
		return
	}

	assert.NotNil(t, actual, msgAndArgs...)
	if actual == nil {
		return
	}

	assert.Equal(t, expected.ID, actual.ID, msgAndArgs...)
	assert.Equal(t, expected.Code, actual.Code, msgAndArgs...)
	assert.Equal(t, expected.Name, actual.Name, msgAndArgs...)
}

func AssertBlockEqual(t *testing.T, expected, actual *models.Block, msgAndArgs ...interface{}) {
	t.Helper()

	if expected == nil && actual == nil {
		return
	}

	assert.NotNil(t, actual, msgAndArgs...)
	if actual == nil {
		return
	}

	assert.Equal(t, expected.ID, actual.ID, msgAndArgs...)
	assert.Equal(t, expected.Name, actual.Name, msgAndArgs...)
	assert.Equal(t, expected.Code, actual.Code, msgAndArgs...)
}
