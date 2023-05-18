package auth

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
	err := errors.New("test error")

	matrix := []struct {
		err    nestedError
		msg    string
		unwrap []error
	}{
		{
			err:    &ErrOTPDisabled{},
			msg:    "otp is not enabled",
			unwrap: []error{},
		},
		{
			err: &ErrFileIO{
				err,
			},
			msg:    "file reading/writing failed",
			unwrap: []error{err},
		},
		{
			err: &ErrFileEncoding{
				err,
			},
			msg:    "file serialisation/deserialisation failed",
			unwrap: []error{err},
		},
		{
			err:    &ErrFileChecksum{},
			msg:    "file checksum validation failed",
			unwrap: []error{},
		},
		{
			err: &ErrFileEncryption{
				err,
			},
			msg:    "file encryption/decryption failed",
			unwrap: []error{err},
		},
	}

	for _, m := range matrix {
		t.Run(fmt.Sprintf("%T", m.err), func(t *testing.T) {
			assert.EqualError(t, m.err, m.msg)
			assert.Equal(t, m.unwrap, m.err.Unwrap())
		})
	}
}
