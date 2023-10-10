package parsers

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type nestedError interface {
	Error() string
	Unwrap() []error
}

func TestErrors(t *testing.T) {
	path := "//div[@id='missing_ref']"

	matrix := []struct {
		err    nestedError
		msg    string
		unwrap []error
	}{
		{
			err: &ErrNotFound{path},
			msg: `element "//div[@id='missing_ref']" not found in document`,
			unwrap: []error{
				&ErrParsing{path, errors.New(`element "//div[@id='missing_ref']" not found in document`)},
			},
		},
		{
			err:    &ErrParsing{path, errors.New(`element "//div[@id='missing_ref']" not found in document`)},
			msg:    `an error happened when parsing "//div[@id='missing_ref']"`,
			unwrap: []error{errors.New(`element "//div[@id='missing_ref']" not found in document`)},
		},
	}

	for _, m := range matrix {
		t.Run(fmt.Sprintf("%T", m.err), func(t *testing.T) {
			assert.EqualError(t, m.err, m.msg)
			assert.Equal(t, m.unwrap, m.err.Unwrap())
		})
	}
}
