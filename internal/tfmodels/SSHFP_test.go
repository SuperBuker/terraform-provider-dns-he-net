package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelSSHFP(t *testing.T) {
	id := uint(1)

	expected := models.SSHFP{
		ID:     &id,
		ZoneID: 1,
		Domain: "example.com",
		TTL:    86400,
		Data:   "4 2 123456789abcdef67890123456789abcdef67890123456789abcdef123456789",
	}

	sshfp := SSHFP{}

	require.NoError(t, sshfp.SetRecord(expected))

	actual, err := sshfp.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
