package parsers

import (
	"errors"
	"fmt"
)

type ErrNotFound struct {
	XPath string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("element \"%s\" not found in document", e.XPath)
}

func (e *ErrNotFound) Unwrap() []error {
	return []error{
		&ErrParsing{e.XPath, errors.New(e.Error())},
	}
}

type ErrParsing struct {
	XPath string
	Err   error
}

func (e *ErrParsing) Error() string {
	return fmt.Sprintf("an error happened when parsing \"%s\" ", e.XPath)
}

func (e *ErrParsing) Unwrap() []error {
	return []error{
		e.Err,
	}
}
