package errors

import (
	"errors"
	"fmt"
)

type DBError struct {
	*BaseError
	Path string
}

func (e DBError) Error() string {
	if e.Path != "" {
		return fmt.Sprintf("database %s failed at path '%s': %s", e.Op, e.Path, e.BaseError.Error())
	}
	return fmt.Sprintf("database %s failed: %s", e.Op, e.BaseError.Error())
}

var (
	ErrDBOpenFailed  = errors.New("database open failed")
	ErrDBPingFailed  = errors.New("database ping failed")
	ErrDBCloseFailed = errors.New("database close failed")
)

func NewDBError(op string, cause error, path ...string) *DBError {
	dbErr := &DBError{
		BaseError: NewBaseError("database", op, "", cause),
	}
	if len(path) > 0 {
		dbErr.Path = path[0]
	}
	return dbErr
}
