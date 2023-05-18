package client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type nestedError interface {
	Error() string
	Unwrap() []error
}

func TestErrors(t *testing.T) {
	matrix := []struct {
		err    nestedError
		msg    string
		unwrap []error
	}{
		{
			err:    &ErrItemNotFound{"resource x"},
			msg:    `item "resource x" not found`,
			unwrap: []error{},
		},
	}

	for _, m := range matrix {
		t.Run(fmt.Sprintf("%T", m.err), func(t *testing.T) {
			assert.EqualError(t, m.err, m.msg)
			assert.Equal(t, m.unwrap, m.err.Unwrap())
		})
	}
}
