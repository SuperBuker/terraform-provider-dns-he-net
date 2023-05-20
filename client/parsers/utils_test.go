package parsers_test

import "fmt"

func errNotFoundString(path string) string {
	return fmt.Sprintf("element %q not found in document", path)
}
