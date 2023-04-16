package client

import "fmt"

// ErrItemNotFound is an error that indicates a resource was not found.
type ErrItemNotFound struct {
	Resource string
}

func (e *ErrItemNotFound) Error() string {
	return fmt.Sprintf(`item "%s" not found`, e.Resource)
}

func (e *ErrItemNotFound) Unwrap() []error {
	return []error{}
}
