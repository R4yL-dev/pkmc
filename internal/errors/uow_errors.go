package errors

import (
	"errors"
	"fmt"
)

type UOWError struct {
	*BaseError
}

func (e UOWError) Error() string {
	return fmt.Sprintf("unit of work %s failed: %s", e.Op, e.BaseError.Error())
}

var (
	ErrUOWBeginFailed  = errors.New("unit of work begin failed")
	ErrUOWCommitFailed = errors.New("unit of work commit failed")
)

func NewUOWError(op string, cause error) *UOWError {
	return &UOWError{
		BaseError: NewBaseError(op, "unit_of_work", "", cause),
	}
}
