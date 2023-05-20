package parsers

import "fmt"

// errParsingNode returns a tailored ErrParsing error.
func errParsingNode(path, field string, err error) error {
	return &ErrParsing{
		fmt.Sprintf("%s // %s", path, field),
		err,
	}
}
