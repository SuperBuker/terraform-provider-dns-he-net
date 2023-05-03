package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelALIAS(t *testing.T) {
	id := uint(1)

	expected := models.ALIAS{
		Id:       &id,
		ParentId: 1,
		Domain:   "example.com",
		TTL:      300,
		Data:     "sub.test.com",
	}

	alias := ALIAS{}

	require.NoError(t, alias.SetRecord(expected))

	actual, err := alias.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
