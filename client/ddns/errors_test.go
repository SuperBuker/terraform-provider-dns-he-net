package ddns

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			err:    &ErrDDNS{"some error"},
			msg:    `ddns update failed: "some error"`,
			unwrap: []error{},
		},
		{
			err:    &ErrAuthFailed{},
			msg:    "authentication failed",
			unwrap: []error{&ErrDDNS{"badauth"}},
		},
		{
			err:    &ErrAbuse{},
			msg:    "authentication blocked due to abuse",
			unwrap: []error{&ErrDDNS{"abuse"}},
		},
		{
			err:    &ErrField{"noipv4"},
			msg:    `missing or malformed field: "myip"`,
			unwrap: []error{&ErrDDNS{"noipv4"}},
		},
		{
			err:    &ErrField{"noipv6"},
			msg:    `missing or malformed field: "myip"`,
			unwrap: []error{&ErrDDNS{"noipv6"}},
		},
		{
			err:    &ErrField{"badip"},
			msg:    `missing or malformed field: "myip"`,
			unwrap: []error{&ErrDDNS{"badip"}},
		},
		{
			err:    &ErrField{"notxt"},
			msg:    `missing or malformed field: "txt"`,
			unwrap: []error{&ErrDDNS{"notxt"}},
		},
		{
			err:    &ErrField{"unknown"},
			msg:    `malformed request: "unknown"`,
			unwrap: []error{&ErrDDNS{"unknown"}},
		},
		{
			err: &ErrAPI{errors.New("some error")},
			msg: "some error",
			unwrap: []error{
				errors.New("some error"),
				&ErrDDNS{"some error"},
			},
		},
		{
			err: &ErrUnknown{"some msg"},
			msg: `unknown error: "some msg"`,
			unwrap: []error{
				&ErrDDNS{"some msg"},
			},
		},
	}

	for _, m := range matrix {
		t.Run(fmt.Sprintf("%T", m.err), func(t *testing.T) {
			require.EqualError(t, m.err, m.msg)
			assert.Equal(t, m.unwrap, m.err.Unwrap())
		})
	}
}
