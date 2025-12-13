package errors

import (
	"errors"
	"fmt"
)

type ServiceError struct {
	*BaseError
	Service string
}

func (e ServiceError) Error() string {
	return fmt.Sprintf("service %s failed: %s", e.Service, e.BaseError.Error())
}

var (
	ErrValidationFailed   = errors.New("service validation failed")
	ErrServiceUnavailable = errors.New("service unavailable")
)

func NewServiceError(op, service, message string, cause error) *ServiceError {
	return &ServiceError{
		BaseError: NewBaseError(op, "service", message, cause),
		Service:   service,
	}
}
