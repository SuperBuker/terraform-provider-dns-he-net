package utils_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
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
			err:    &utils.ErrNotImplemented{},
			msg:    "feature is not implemented",
			unwrap: []error{},
		},
		{
			err:    utils.NewErrCasting("some string", 0),
			msg:    "type casting failed: expected string, got int",
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
