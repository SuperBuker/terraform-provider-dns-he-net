package models

import (
	"fmt"
)

type ErrFormat struct {
	Field  string
	Reason string
}

func (e *ErrFormat) Error() string {
	if e.Field == "" {
		return fmt.Sprintf("invalid record, %s", e.Reason)
	} else {
		return fmt.Sprintf("invalid \"%s\", %s", e.Field, e.Reason)
	}

}
