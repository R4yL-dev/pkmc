package database

import (
	"errors"
	"os"
	"testing"

	customErr "github.com/R4yL-dev/pkmc/internal/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitDB_Success(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_db_*.db")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	db, err := InitDB(tmpFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, db)
	defer CloseDB(db)
}

func TestInitDB_InvalidPath(t *testing.T) {
	_, err := InitDB("/invalid/path/to/db.db")
	assert.Error(t, err)

	var dbErr *customErr.DBError
	assert.True(t, errors.As(err, &dbErr), "expected DBError")
	assert.Equal(t, "open", dbErr.Op)
	assert.Equal(t, "database", dbErr.Domain)
	assert.Equal(t, "/invalid/path/to/db.db", dbErr.Path)
	assert.NotNil(t, dbErr.Cause)
}

func TestInitDB_GetSQLDBFailure(t *testing.T) {
	t.Skip("GetSQLDB failure hard to simulate without mocks")
}

func TestInitDB_PingFailure(t *testing.T) {
	t.Skip("Ping failure hard to simulate without mocks")
}

func TestCloseDB_Success(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_db_*.db")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	db, err := InitDB(tmpFile.Name())
	require.NoError(t, err)

	err = CloseDB(db)
	assert.NoError(t, err)
}

func TestCloseDB_NilDB(t *testing.T) {
	err := CloseDB(nil)
	assert.NoError(t, err)
}

func TestCloseDB_AlreadyClosed(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_db_*.db")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	db, err := InitDB(tmpFile.Name())
	require.NoError(t, err)

	CloseDB(db)

	err = CloseDB(db)
	if err != nil {
		var dbErr *customErr.DBError
		assert.True(t, errors.As(err, &dbErr), "expected DBError")
		assert.Equal(t, "close", dbErr.Op)
		assert.NotNil(t, dbErr.Cause)
	} else {
		t.Log("Closing already closed DB succeeded (expected behavior)")
	}
}

func TestCloseDB_GetSQLDBFailure(t *testing.T) {
	t.Skip("GetSQLDB failure hard to simulate without mocks")
}
