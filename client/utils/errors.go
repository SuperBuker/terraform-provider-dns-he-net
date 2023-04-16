package utils

import "fmt"

// ErrNotImplemented is an error that indicates a feature is not implemented.
type ErrNotImplemented struct{}

func (e *ErrNotImplemented) Error() string {
	return "feature is not implemented"
}

func (e *ErrNotImplemented) Unwrap() []error {
	return []error{}
}

// ErrCasting is an error that indicates a type casting failed.
type ErrCasting struct {
	ExpectedType string
	ActualType   string
}

func (e *ErrCasting) Error() string {
	return fmt.Sprintf("type casting failed: expected %s, got %s", e.ExpectedType, e.ActualType)
}

func (e *ErrCasting) Unwrap() []error {
	return []error{}
}

func NewErrCasting(expected, actual interface{}) *ErrCasting {
	return &ErrCasting{
		ExpectedType: fmt.Sprintf("%T", expected),
		ActualType:   fmt.Sprintf("%T", actual),
	}
}
