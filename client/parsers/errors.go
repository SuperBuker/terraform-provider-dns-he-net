package parsers

import (
	"errors"
	"fmt"
)

// ErrNotFound is returned when the element is not found in the document.
type ErrNotFound struct {
	XPath string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("element %q not found in document", e.XPath)
}

func (e *ErrNotFound) Unwrap() []error {
	return []error{
		&ErrParsing{e.XPath, errors.New(e.Error())},
	}
}

// ErrParsing is returned when an error happens when parsing the document.
type ErrParsing struct {
	XPath string
	Err   error
}

func (e *ErrParsing) Error() string {
	return fmt.Sprintf("an error happened when parsing %q", e.XPath)
}

func (e *ErrParsing) Unwrap() []error {
	return []error{
		e.Err,
	}
}
