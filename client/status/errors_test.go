package status

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ error = (*ErrNoAuth)(nil)
var _ error = (*ErrPartialAuth)(nil)
var _ error = (*ErrAuthFailed)(nil)
var _ error = (*ErrMissingOTPAuth)(nil)
var _ error = (*ErrOTPAuthFailed)(nil)
var _ error = (*ErrUnknownAuth)(nil)
var _ error = (*ErrHeNet)(nil)

func TestErrorMessagesAndUnwrap(t *testing.T) {
	t.Run("ErrNoAuth", func(t *testing.T) {
		e := &ErrNoAuth{}
		assert.Equal(t, "not authenticated", e.Error())
		u := e.Unwrap()
		if assert.Len(t, u, 1) {
			_, ok := u[0].(*ErrHeNet)
			assert.True(t, ok)
		}
	})

	t.Run("ErrPartialAuth", func(t *testing.T) {
		e := &ErrPartialAuth{}
		assert.Equal(t, "authentication not completed", e.Error())
		u := e.Unwrap()
		if assert.Len(t, u, 2) {
			_, ok1 := u[0].(*ErrNoAuth)
			_, ok2 := u[1].(*ErrHeNet)
			assert.True(t, ok1 && ok2)
		}
	})

	t.Run("ErrAuthFailed", func(t *testing.T) {
		e := &ErrAuthFailed{error: "bad creds"}
		assert.Equal(t, "authentication failed: bad creds", e.Error())
		u := e.Unwrap()
		if assert.Len(t, u, 2) {
			_, ok1 := u[0].(*ErrNoAuth)
			he, ok2 := u[1].(*ErrHeNet)
			assert.True(t, ok1 && ok2)
			if ok2 {
				assert.Equal(t, "bad creds", he.error)
			}
		}
	})

	t.Run("ErrMissingOTPAuth", func(t *testing.T) {
		e := &ErrMissingOTPAuth{}
		assert.Equal(t, "missing OTP authentication", e.Error())
		u := e.Unwrap()
		if assert.Len(t, u, 2) {
			_, ok1 := u[0].(*ErrPartialAuth)
			_, ok2 := u[1].(*ErrHeNet)
			assert.True(t, ok1 && ok2)
		}
	})

	t.Run("ErrOTPAuthFailed", func(t *testing.T) {
		e := &ErrOTPAuthFailed{error: "totp bad"}
		assert.Equal(t, "authentication failed: totp bad", e.Error())
		u := e.Unwrap()
		if assert.Len(t, u, 2) {
			authFail, ok1 := u[0].(*ErrAuthFailed)
			he, ok2 := u[1].(*ErrHeNet)
			assert.True(t, ok1 && ok2)
			if ok1 {
				assert.Equal(t, "totp bad", authFail.error)
			}
			if ok2 {
				assert.Equal(t, "totp bad", he.error)
			}
		}
	})

	t.Run("ErrUnknownAuth", func(t *testing.T) {
		e := &ErrUnknownAuth{}
		assert.Equal(t, "unknown authentication status", e.Error())
		u := e.Unwrap()
		if assert.Len(t, u, 2) {
			_, ok1 := u[0].(*ErrAuthFailed)
			_, ok2 := u[1].(*ErrHeNet)
			assert.True(t, ok1 && ok2)
		}
	})

	t.Run("ErrHeNet", func(t *testing.T) {
		e := &ErrHeNet{error: "some he.net error"}
		assert.Equal(t, "some he.net error", e.Error())
	})
}

func TestErrorScoreValues(t *testing.T) {
	assert.Equal(t, 7, errorScore(&ErrOTPAuthFailed{}))
	assert.Equal(t, 6, errorScore(&ErrAuthFailed{}))
	assert.Equal(t, 5, errorScore(&ErrMissingOTPAuth{}))
	assert.Equal(t, 4, errorScore(&ErrPartialAuth{}))
	assert.Equal(t, 3, errorScore(&ErrNoAuth{}))
	assert.Equal(t, 2, errorScore(&ErrUnknownAuth{}))
	assert.Equal(t, 1, errorScore(&ErrHeNet{}))
	assert.Equal(t, 0, errorScore(nil))
}
