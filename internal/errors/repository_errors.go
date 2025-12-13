package errors

import (
	"errors"
	"fmt"
)

type RepositoryError struct {
	*BaseError
	Entity string
	Key    string
}

func (e RepositoryError) Error() string {
	return fmt.Sprintf("repository %s failed for %s '%s': %s", e.Entity, e.Op, e.Key, e.BaseError.Error())
}

var (
	ErrEntityNotFound      = errors.New("entity not found")
	ErrConstraintViolation = errors.New("constraint violation")
)

func NewRepositoryError(op, entity, key string, cause error) *RepositoryError {
	return &RepositoryError{
		BaseError: NewBaseError(op, "repository", "", cause),
		Entity:    entity,
		Key:       key,
	}
}
