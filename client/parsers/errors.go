package parsers

import "fmt"

type ErrNotFound struct {
	XPath string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("element \"%s\" not found in document", e.XPath)
}
