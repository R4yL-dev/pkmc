package errors

import "fmt"

type BaseError struct {
	Op      string
	Domain  string
	Message string
	Cause   error
}

func (e BaseError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s %s failed: %s (%v)", e.Domain, e.Op, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s %s failed: %s", e.Domain, e.Op, e.Message)
}

func (e BaseError) Unwrap() error {
	return e.Cause
}

func NewBaseError(op, domain, message string, cause error) *BaseError {
	return &BaseError{
		Op:      op,
		Domain:  domain,
		Message: message,
		Cause:   cause,
	}
}
