package errors

import (
	"errors"
	"fmt"
)

type UOWError struct {
	*BaseError
	Operation string
}

func (e UOWError) Error() string {
	return fmt.Sprintf("unit of work %s failed: %s", e.Operation, e.BaseError.Error())
}

var (
	ErrUOWBeginFailed  = errors.New("unit of work begin failed")
	ErrUOWCommitFailed = errors.New("unit of work commit failed")
)

func NewUOWError(operation string, cause error) *UOWError {
	return &UOWError{
		BaseError: NewBaseError(operation, "unit_of_work", "", cause),
		Operation: operation,
	}
}
