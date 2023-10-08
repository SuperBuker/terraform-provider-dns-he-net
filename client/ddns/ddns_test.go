package ddns

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDDNS(t *testing.T) {
	domains := generateSubDomains("hostname-%04d.example.com", 10000, 3)

	t.Run("CheckAuth", func(t *testing.T) {
		ok, err := CheckAuth(context.Background(), domains[0], "password")

		require.NoError(t, err)
		assert.False(t, ok)
	})

	t.Run("UpdateIP", func(t *testing.T) {
		ok, err := UpdateIP(context.Background(), domains[1], "password", "0.0.0.0")

		require.ErrorIs(t, err, &ErrAuthFailed{})
		assert.False(t, ok)
	})

	t.Run("UpdateTXT", func(t *testing.T) {
		ok, err := UpdateTXT(context.Background(), domains[2], "password", "some text")

		require.ErrorIs(t, err, &ErrAuthFailed{})
		assert.False(t, ok)
	})
}

func TestProcessResponse(t *testing.T) {
	matrix := []struct {
		response string
		ok       bool
		error    error
	}{
		{
			response: "good 0.0.0.0",
			ok:       true,
			error:    nil,
		},
		{
			response: "nochg",
			ok:       false,
			error:    nil,
		},
		{
			response: "badauth",
			ok:       false,
			error:    &ErrAuthFailed{},
		},
		{
			response: "abuse",
			ok:       false,
			error:    &ErrAbuse{},
		},
		{
			response: "badip",
			ok:       false,
			error:    &ErrField{"badip"},
		},
		{
			response: "another error",
			ok:       false,
			error:    &ErrUnknown{"another error"},
		},
	}

	for _, m := range matrix {
		t.Run(fmt.Sprintf("%q", m.response), func(t *testing.T) {
			ok, err := processResponse(m.response)
			require.Equal(t, m.error, err)
			assert.Equal(t, m.ok, ok)
		})
	}
}

// This test needs to be extended to validate other error cases.
